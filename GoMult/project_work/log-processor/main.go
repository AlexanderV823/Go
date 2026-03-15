package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const filename string = "./log-processor/testdata/log.csv"
	const numWorkers int = 3
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	
	go func() {
		<-sigCh
		cancel()
	}()

	// Получаем канал в который пишутся строки лога
	entryCh, err := readLogs(ctx, filename)
	if err != nil {
		log.Printf("error read CSV: %v", err)
	}

	// Передаем канал со строками лога в worker pool
	processedCh := processLogs(ctx, entryCh, numWorkers)
}
