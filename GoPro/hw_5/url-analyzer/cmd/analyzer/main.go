package main

import (
	"fmt"
	"sort"
	"time"
	"url-analyzer/internal/pool"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://habr.com",
		"https://yandex.ru",
		"https://vk.com",
		"https://reddit.com",
		"https://youtube.com",
	}

	jobs := make([]pool.Job, len(urls))
	for i, url := range urls {
		jobs[i] = pool.Job{ID: i + 1, URL: url}
	}

	startTime := time.Now()

	// Инициализируем пул воркеров
	p := pool.NewPool(3)
	results := p.Start(jobs)

	totalTime := time.Since(startTime)

	// Сортируем результаты по JobID для стабильного вывода отчета
	sort.Slice(results, func(i, j int) bool {
		return results[i].JobID < results[j].JobID
	})

	// Вывод финального отчета
	fmt.Println("=== ФИНАЛЬНЫЙ ОТЧЁТ ПО ОБРАБОТКЕ URL ===")
	fmt.Println("ID | URL | Статус | Время")
	fmt.Println("------------------------------------------------------------")

	var totalDuration time.Duration
	successCount := 0

	for _, res := range results {
		fmt.Printf("%d | %s | %s | %v\n", res.JobID, res.URL, res.Status, res.Duration)
		totalDuration += res.Duration
		if res.Status == "Успешно" {
			successCount++
		}
	}

	avgDuration := totalDuration / time.Duration(len(results))

	fmt.Println("------------------------------------------------------------")
	fmt.Println("=== Общая статистика ===")
	fmt.Printf("Количество успешных операций: %d из %d\n", successCount, len(urls))
	fmt.Printf("Общее время работы всех имитаций: %.2fs\n", totalTime.Seconds())
	fmt.Printf("Среднее время выполнения одного запроса: %v\n", avgDuration.Round(time.Millisecond))
}
