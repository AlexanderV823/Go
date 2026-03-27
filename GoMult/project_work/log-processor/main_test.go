package main

import (
	"context"
	"os"
	"testing"
)

func TestLogPipeline(t *testing.T) {
	// 1. Подготовка тестовых данных
	content := "timestamp,ip,method,url,status,duration\n" +
		"2024-01-15 10:30:00,192.168.1.100,GET,/api/users,200,150\n" +
		"2024-01-15 10:30:01,192.168.1.101,POST,/api/users,201,200\n" +
		"2024-01-15 10:30:02,192.168.1.100,GET,/api/users/123,404,50\n" +
		"2024-01-15 10:30:03,192.168.1.102,GET,/api/products,500,1500\n" +
		"2024-01-15 10:30:04,192.168.1.100,GET,/api/orders,200,100\n"

	tmpFile, err := os.CreateTemp("", "test_log_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // удаляем после теста

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	minStatus := 400
	numWorkers := 3

	pipeCh1, err := readLogs(ctx, tmpFile.Name())
	if err != nil {
		t.Fatalf("readLogs error: %v", err)
	}

	pipeCh2 := processLogs(ctx, pipeCh1, numWorkers)

	statCh, filterCh := fanOut(ctx, pipeCh2)

	var result []LogEntry
	for entry := range filterLogs(ctx, filterCh, minStatus) {
		result = append(result, entry)
	}
	if len(result) != 2 {
		t.Errorf("Ожидалось 2 записи с ошибками, получили %d", len(result))
	}

	// Передаем канал на финальный сбор статистики:
	stat := calculateStats(ctx, statCh)

	expectedTotal := 5
	if stat.TotalRequests != expectedTotal {
		t.Errorf("Ожидалось %d запросов после фильтра, получили %d", expectedTotal, stat.TotalRequests)
	}

	if stat.ErrorCount != 2 {
		t.Errorf("Ожидалось 2 ошибки, получили %d", stat.ErrorCount)
	}

	count := stat.RequestsByIP["192.168.1.100"]
	if count != 3 {
		t.Errorf("Для IP 192.168.1.100 ожидалось 3 запроса, получено %d", count)
	}

	expectedAvg := 400.00
	if stat.AverageRespTime != expectedAvg {
		t.Errorf("Ожидалось среднее время %.2f, получили %.2f", expectedAvg, stat.AverageRespTime)
	}
}
