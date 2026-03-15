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
	
}

// filterLogs - фильтрует записи по статус-коду
func filterLogs(input <-chan LogEntry, minStatus int) <-chan LogEntry {

}

// calculateStats - выполняет подсчет статистики
func calculateStats(input <-chan LogEntry) Statistics {

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

}
