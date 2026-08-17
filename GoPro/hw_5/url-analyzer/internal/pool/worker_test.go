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
	numWorkers := 2

	// Запуск тестируемой функции
	results := RunPool(testURLs, numWorkers)

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
			t.Errorf("Для URL %s получен некорректный статус: %s", res.Job.URL, res.Status)
		}

		// Проверяем, что время выполнения больше нуля
		if res.Duration <= 0 {
			t.Errorf("Для URL %s зафиксировано некорректное время: %v", res.Job.URL, res.Duration)
		}

		// Отмечаем, что данный URL был обработан
		if _, exists := urlCheckMap[res.Job.URL]; exists {
			urlCheckMap[res.Job.URL] = true
		} else {
			t.Errorf("Получен результат для неизвестного URL: %s", res.Job.URL)
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
	var testURLs []string
	numWorkers := 3

	results := RunPool(testURLs, numWorkers)

	if len(results) != 0 {
		t.Errorf("Ожидался пустой результат для пустого списка URL, получено элементов: %d", len(results))
	}
}
