package pool

import (
	"testing"
)

// TestRunPool проверяет базовую логику работы Worker Pool
func TestRunPool(t *testing.T) {
	// Входные тестовые данные
	testURLs := []string{
		"https://test1.com",
		"https://test2.com",
		"https://test3.com",
	}

	// Формируем срез задач (Job) на основе URL
	jobs := make([]Job, len(testURLs))
	for i, url := range testURLs {
		jobs[i] = Job{ID: i + 1, URL: url}
	}

	numWorkers := 2

	// Инициализируем и запускаем пул воркеров
	pool := NewPool(numWorkers)
	results := pool.Start(jobs)

	// Проверка 1: Количество результатов должно совпадать с количеством задач
	if len(results) != len(testURLs) {
		t.Errorf("Ожидалось %d результатов, но получено %d", len(testURLs), len(results))
	}

	// Создаем карту для быстрой проверки наличия всех URL в результатах
	urlCheckMap := make(map[string]bool)
	for _, url := range testURLs {
		urlCheckMap[url] = false
	}

	// Проверка 2: Проверяем корректность данных в результатах
	for _, res := range results {
		// Проверяем, что статус успешный
		if res.Status != "Успешно" {
			t.Errorf("Для ID %d получен некорректный статус: %s", res.JobID, res.Status)
		}

		// Проверяем, что время выполнения больше нуля
		if res.Duration <= 0 {
			t.Errorf("Для URL %s зафиксировано некорректное время: %v", res.URL, res.Duration)
		}

		// Отмечаем, что данный URL был обработан
		if _, exists := urlCheckMap[res.URL]; exists {
			urlCheckMap[res.URL] = true
		} else {
			t.Errorf("Получен результат для неизвестного URL: %s", res.URL)
		}
	}

	// Проверка 3: Убеждаемся, что ни один URL не был пропущен
	for url, processed := range urlCheckMap {
		if !processed {
			t.Errorf("URL %s не был обработан воркерами", url)
		}
	}
}

// TestRunPool_EmptyList проверяет поведение пула при пустом списке задач
func TestRunPool_EmptyList(t *testing.T) {
	var jobs []Job
	numWorkers := 3

	pool := NewPool(numWorkers)
	results := pool.Start(jobs)

	if len(results) != 0 {
		t.Errorf("Ожидался пустой результат для пустого списка URL, получено элементов: %d", len(results))
	}
}

// TestRunPool_ZeroWorkers проверяет защиту от передачи 0 воркеров
func TestRunPool_ZeroWorkers(t *testing.T) {
	jobs := []Job{
		{ID: 1, URL: "https://test1.com"},
	}

	// Передаем 0 воркеров (конструктор должен безопасно переключить на 1)
	pool := NewPool(0)
	results := pool.Start(jobs)

	if len(results) != 1 {
		t.Errorf("Ожидался 1 результат при автоматическом исправлении пула на 1 воркера, получено: %d", len(results))
	}
}
