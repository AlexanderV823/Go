package pool

import (
	"math/rand"
	"sync"
	"time"
)

// worker обрабатывает задачи из канала jobs и отправляет результаты в results.
func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Симулируем случайный пинг/задержку от 100 до 600 миллисекунд
		randomDuration := time.Duration(100+rand.Intn(500)) * time.Millisecond
		time.Sleep(randomDuration)

		// Отправляем успешный результат выполнения задачи
		results <- Result{
			Job:      job,
			Status:   "Успешно",
			Duration: randomDuration,
		}
	}
}

// RunPool — точка входа в пакет. Управляет паттернами Fan-out и Fan-in.
func RunPool(urls []string, numWorkers int) []Result {
	// Создаем каналы с буфером под размер переданных URL
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	// [Fan-out] Инициализируем пул воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Генерация задач и отправка в очередь
	for i, url := range urls {
		jobs <- Job{
			ID:  i + 1,
			URL: url,
		}
	}
	close(jobs) // Сигнал воркерам, что новых задач больше не будет

	// [Fan-in] Ожидание завершения всех горутин и закрытие канала результатов
	go func() {
		wg.Wait()
		close(results)
	}()

	// Агрегируем результаты из канала в единый слайс
	var finalReport []Result
	for res := range results {
		finalReport = append(finalReport, res)
	}

	return finalReport
}
