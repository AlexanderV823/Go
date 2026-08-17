package main

import (
	"fmt"
	"time"

	// Подключаем наш внутренний пакет.
	// "url-analyzer" — это имя модуля, которое вы указали при `go mod init`
	"url-analyzer/internal/pool"
)

func main() {
	// Исходный массив адресов для проверки производительности
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://yandex.ru",
		"https://golang.org",
		"https://habr.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://docker.com",
	}

	// Ограничиваем одновременные запросы количеством в 5 воркеров
	numWorkers := 5

	// Передаем работу изолированному модулю пула
	finalReport := pool.RunPool(urls, numWorkers)

	// Форматированный вывод таблицы результатов
	fmt.Println("=== ФИНАЛЬНЫЙ ОТЧЁТ ПО ОБРАБОТКЕ URL ===")
	fmt.Printf("%-3s | %-26s | %-10s | %-10s\n", "ID", "URL", "Статус", "Время")
	fmt.Println("------------------------------------------------------------")

	var totalDuration time.Duration
	for _, res := range finalReport {
		fmt.Printf("%-3d | %-26s | %-10s | %v\n", res.Job.ID, res.Job.URL, res.Status, res.Duration)
		totalDuration += res.Duration
	}

	// Вычисление и вывод общих метрик
	fmt.Println("------------------------------------------------------------")
	fmt.Println("=== Общая статистика ===")
	totalOps := len(finalReport)
	fmt.Printf("Количество успешных операций: %d из %d\n", totalOps, len(urls))

	if totalOps > 0 {
		avgDuration := totalDuration / time.Duration(totalOps)
		fmt.Printf("Общее время работы всех имитаций: %v\n", totalDuration)
		fmt.Printf("Среднее время выполнения одного запроса: %v\n", avgDuration)
	}
}
