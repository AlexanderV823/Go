package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type LogEntry struct {
	Timestamp    string // время в формате "2024-01-15 10:30:00"
	IP           string // IP адрес клиента
	Method       string // HTTP метод (GET, POST и т.д.)
	URL          string // путь запроса
	StatusCode   int    // HTTP статус код
	ResponseTime int    // время ответа в миллисекундах
}

type Statistics struct {
	TotalRequests   int            // общее количество запросов
	ErrorCount      int            // количество ошибок (статус >= 400)
	RequestsByIP    map[string]int // количество запросов с каждого IP
	AverageRespTime float64        // среднее время ответа
}

func main() {

	numWorkers := flag.Int("workers", 3, "количество воркеров для обработки, значение по умолчанию 3")
	minStatus := flag.Int("status", 400, "минимальный статус-код для фильтрации, значение по умолчанию 400")
	topIP := flag.Int("top", 3, "количество топ IP-адресов для вывода, значение по умолчанию 3")
	filename := flag.String("logfile", "testdata/log.csv", "путь к файлу лога, значение по умолчанию testdata/log.csv")

	flag.Parse()

	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		log.Fatalf("error: file '%s' not found", *filename)
	}

	fmt.Printf("Запуск: файл=%s, воркеры=%d, мин. статус=%d ТОП-%d IP-адресов\n", *filename, *numWorkers, *minStatus, *topIP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		log.Printf("prosecc canceled")
		cancel()
	}()

	// Получаем канал в который пишутся строки лога:
	pipeCh1, err := readLogs(ctx, *filename)
	if err != nil {
		log.Printf("error read CSV: %v", err)
	}

	// Передаем канал со строками лога в worker pool:
	pipeCh2 := processLogs(ctx, pipeCh1, *numWorkers)

	// Передаем на фильтрацию канал
	pipeCh3 := filterLogs(ctx, pipeCh2, *minStatus)

	// Передаем канал на финальный сбор статистики:
	stat := calculateStats(ctx, pipeCh3)

	printStatistics(stat)

	printTopIPs(stat.RequestsByIP, *topIP)
}
