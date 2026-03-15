package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

// readLogs - выполняет чтение логов из файлов
func readLogs(ctx context.Context, filename string) (<-chan LogEntry, error) {
	outCh := make(chan LogEntry)

	go func() {
		defer close(outCh)

		file, err := os.Open(filename)
		if err != nil {
			log.Printf("open CSV file error: %v", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		_, _ = reader.Read() // Skip top line

		for {
			select {
			case <-ctx.Done():
				log.Printf("context cancel: %v", ctx.Err())
				return
			default:
				line, err := reader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Printf("string read error: %v", err)
					continue
				}

				logEntry, err := parseLogLine(line)
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
	outCh := make(chan LogEntry)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case entry, ok := <-input:
					if !ok {
						log.Printf("read chan logEntry error")
						return
					}
					select {
					case <-ctx.Done():
						log.Printf("context canceled: %v", ctx.Err())
						return
					case outCh <- entry:
					}
				}
			}
		}()
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
		select {
		case <-ctx.Done():
			log.Printf("context cancel: %v", ctx.Err())
			return
		case entry, ok := <-input:
			if !ok {
				log.Printf("read chan logEntry error")
				return
			}
			if entry.StatusCode >= minStatus {
				select {
				case <-ctx.Done():
					log.Printf("context cancel: %v", ctx.Err())
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

	select {
	case <-ctx.Done():
		log.Printf("context cancel: %v", ctx.Err())
		return Statistics{}
	case entry, ok := <-input:
		if !ok {
			log.Printf("read chan logEntry error")
			return Statistics{}
		}
		// Наполняем возвращаемую структуру результатами чтения из исходного файла:
		// 1. Перезаписывем счетчик общего числа запросов
		stats.TotalRequests++
		// 2. Проверяем код ошибки и при условии пополняем счетчик ошибок
		if entry.StatusCode >= 400 {
			stats.ErrorCount++
		}
		// 3. Пополняем счетчики числа запросов с разных IP
		stats.RequestsByIP[entry.IP]++
		// 4.1. Накапливаем общее время запросов
		totalResponseTime += entry.ResponseTime
	}
	if stats.TotalRequests > 0 {
		// 4.2. Записываем среднее время запросов
		stats.AverageRespTime = float64(totalResponseTime) / float64(stats.TotalRequests)
	}
	return stats
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

// printTopIPs - выводит топ IP-адресов
func printTopIPs(requestsByIP map[string]int, n int) {
	tempMap := make(map[string]int, len(requestsByIP))

	for key, value := range requestsByIP {
		tempMap[key] = value
	}

	fmt.Printf("\nТоп-%d IP-адресов:\n", n)

	for i := 1; i <= n; i++ {
		maxCount := -1
		maxIP := ""

		for ip, count := range tempMap {
			if count > maxCount {
				maxCount = count
				maxIP = ip
			}
		}

		if maxIP == "" {
			break
		}

		fmt.Printf("%d. %-15s — %d запросов\n", i, maxIP, maxCount)

		delete(tempMap, maxIP)
	}
}
