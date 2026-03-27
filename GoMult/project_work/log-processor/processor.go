package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

type IPStat struct {
	IP    string
	Count int
}

// readLogs - выполняет чтение логов из файлов
func readLogs(ctx context.Context, filename string) (<-chan LogEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open CSV file: %w", err)
	}

	outCh := make(chan LogEntry, 100)

	go func() {
		defer close(outCh)
		defer file.Close()
		reader := csv.NewReader(file)
		// Пропуск заголовка
		if _, err := reader.Read(); err != nil {
			if err != io.EOF {
				log.Printf("header CSV file error: %v", err)
			}
			return
		}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, err := reader.Read()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Printf("string read error: %v", err)
					continue
				}

				logEntry, err := parseLogLine(line) // Перенести бы эту обработку в worker pool, а из этой функции возвращать только прочитанную строку.
				// Но по условимя задачи readLogs должна возвращать (<-chan LogEntry, error)
				if err != nil {
					log.Printf("pass line: %v", err)
					continue
				}

				select {
				case <-ctx.Done():
					return
				case outCh <- logEntry:
				}
			}
		}
	}()
	return outCh, nil
}

// processLogs - обрабатывает логи через worker pool
func processLogs(ctx context.Context, input <-chan LogEntry, numWorkers int) <-chan LogEntry {
	if numWorkers < 3 {
		numWorkers = 3
	}

	outCh := make(chan LogEntry)
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case entry, ok := <-input:
				if !ok {
					return
				}
				// Здесь могла бы быть обработка, но по условиям задачи нужно передавать и возвращать один канал <-chan LogEntry
				select {
				case <-ctx.Done():
					return
				case outCh <- entry:
				}
			}
		}
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}
	go func() {
		wg.Wait()
		close(outCh)
	}()
	return outCh
}

// filterLogs - фильтрует записи по статус-коду
func filterLogs(ctx context.Context, input <-chan LogEntry, minStatus int) <-chan LogEntry {
	outCh := make(chan LogEntry)

	go func() {
		defer close(outCh)
		for {
			select {
			case <-ctx.Done():
				return
			case entry, ok := <-input:
				if !ok {
					return
				}
				if entry.StatusCode < minStatus {
					continue
				}
				select {
				case <-ctx.Done():
					return
				case outCh <- entry:
				}
			}
		}
	}()
	return outCh
}

// calculateStats - выполняет подсчет статистики
func calculateStats(ctx context.Context, input <-chan LogEntry) Statistics {
	stats := Statistics{
		RequestsByIP: make(map[string]int)}
	var totalResponseTime int

	for {
		select {
		case <-ctx.Done():
			log.Printf("context cancel: %v", ctx.Err())
			return stats
		case entry, ok := <-input:
			if !ok {
				if stats.TotalRequests > 0 {
					stats.AverageRespTime = float64(totalResponseTime) / float64(stats.TotalRequests)
				}
				return stats
			}

			stats.TotalRequests++
			if entry.StatusCode >= 400 {
				stats.ErrorCount++
			}
			stats.RequestsByIP[entry.IP]++
			totalResponseTime += entry.ResponseTime
		}
	}
}

// parseLogLine - выполняет чтение строки CSV в структуру LogEntry
func parseLogLine(line []string) (LogEntry, error) {

	if len(line) < 6 {
		return LogEntry{}, fmt.Errorf("not enough data in the row: %v", line)
	}

	statusCode, err := strconv.Atoi(line[4])
	if err != nil {
		return LogEntry{}, fmt.Errorf("status code convert error: %v", err)
	}

	responseTime, err := strconv.Atoi(line[5])
	if err != nil {
		return LogEntry{}, fmt.Errorf("response time convert error: %v", err)
	}

	return LogEntry{
			Timestamp:    line[0],
			IP:           line[1],
			Method:       line[2],
			URL:          line[3],
			StatusCode:   statusCode,
			ResponseTime: responseTime},
		nil
}

func fanOut(ctx context.Context, input <-chan LogEntry) (<-chan LogEntry, <-chan LogEntry) {
	out1 := make(chan LogEntry, 1000)
	out2 := make(chan LogEntry, 1000)

	go func() {
		defer close(out1)
		defer close(out2)
		for entry := range input {
			select {
			case <-ctx.Done():
				return
			case out1 <- entry:
			}
			select {
			case <-ctx.Done():
				return
			case out2 <- entry:
			}
		}
	}()
	return out1, out2
}

// printTopIPs - выводит топ (n) IP-адресов
func printTopIPs(requestsByIP map[string]int, n int) {
	if len(requestsByIP) == 0 {
		fmt.Println("\nДанные по IP отсутствуют.")
		return
	}

	stats := make([]IPStat, 0, len(requestsByIP))
	for ip, count := range requestsByIP {
		stats = append(stats, IPStat{IP: ip, Count: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})

	fmt.Printf("\nТоп-%d IP-адресов:\n", n)
	for i := 0; i < n && i < len(stats); i++ {
		fmt.Printf("%d. %-15s — %d запрос(ов)\n", i+1, stats[i].IP, stats[i].Count)
	}
}

// printStatistics - выводит расчитанную статистику
func printStatistics(stats Statistics) {

	var errorPercent float64

	if stats.TotalRequests > 0 {
		errorPercent = float64(stats.ErrorCount) / float64(stats.TotalRequests) * 100
	} else {
		errorPercent = 0
	}

	fmt.Printf("Всего запросов:      %d\n", stats.TotalRequests)
	fmt.Printf("Количество ошибок:   %d (%.2f%%)\n",
		stats.ErrorCount,
		errorPercent,
	)
	fmt.Printf("Среднее время ответа: %.2f мс\n", stats.AverageRespTime)

}
