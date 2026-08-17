package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job представляет задачу для обработки.
type Job struct {
	ID  int
	URL string
}

// Result представляет результат обработки задачи.
type Result struct {
	Job      Job
	Status   string
	Duration time.Duration
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Цикл for range по каналу с заданиями
	for job := range jobs {
		// Имитация случайной задержки от 100 до 600 миллисекунд
		randomDuration := time.Duration(100+rand.Intn(500)) * time.Millisecond
		time.Sleep(randomDuration)

		// Отправка результата обработки в канал результатов
		results <- Result{
			Job:      job,
			Status:   "Успешно",
			Duration: randomDuration,
		}
	}
}

func main() {
	// Исходный список URL-адресов
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

	numWorkers := 5 // Фиксированное количество горутин-воркеров

	// Каналы буферизированы размером списка URL, чтобы избежать блокировок
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	// Запуск фиксированного количества воркеров (Fan-out)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Наполняем канал jobs заданиями на основе списка URL
	for i, url := range urls {
		jobs <- Job{
			ID:  i + 1,
			URL: url,
		}
	}
	// Закрываем канал, чтобы воркеры знали, когда остановиться
	close(jobs)

	// Отдельная горутина ждет завершения воркеров и закрывает канал результатов
	go func() {
		wg.Wait()
		close(results)
	}()

	var finalReport []Result
	var totalDuration time.Duration

	// Чтение всех результатов в главной горутине через for range
	for res := range results {
		finalReport = append(finalReport, res)
		totalDuration += res.Duration
	}

	// Вывод финального отчёта
	fmt.Println("=== ФИНАЛЬНЫЙ ОТЧЁТ ПО ОБРАБОТКЕ URL ===")
	fmt.Printf("%-3s | %-26s | %-10s | %-10s\n", "ID", "URL", "Статус", "Время")
	fmt.Println("------------------------------------------------------------")

	for _, res := range finalReport {
		fmt.Printf("%-3d | %-26s | %-10s | %v\n", res.Job.ID, res.Job.URL, res.Status, res.Duration)
	}

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