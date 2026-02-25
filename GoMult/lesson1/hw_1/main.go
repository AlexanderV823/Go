package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task представляет задачу для выполнения
type Task struct {
	ID   int
	Name string
	Type string
}

// Result представляет результат выполнения задачи
type Result struct {
	TaskID   int
	TaskName string
	Success  bool
	Duration time.Duration
	Message  string
}

// simulateIOWork имитирует I/O операцию (чтение файла, сетевой запрос)
func simulateIOWork(task Task, wg *sync.WaitGroup, results chan<- Result) {
	// TODO: Реализуйте функцию для имитации I/O работы
	// Подсказки:
	// 1. Не забудьте вызвать wg.Done() в конце функции (используйте defer)
	// 2. Выведите сообщение о начале работы
	// 3. Используйте time.Sleep() для имитации I/O операции (1-3 секунды)
	// 4. Создайте объект Result и отправьте его в канал results
	// 5. Выведите сообщение о завершении работы

	defer wg.Done()

	start := time.Now()

	fmt.Printf("🔄 Начинаю I/O задачу %d: %s\n", task.ID, task.Name)

	time.Sleep(3 * time.Second)

	elapsed := time.Since(start)

	res := Result{
		TaskID:   task.ID,
		TaskName: task.Name,
		Success:  true,
		Duration: elapsed,
		Message:  "I/O данные обработаны"}

	results <- res

	fmt.Printf("🔄 Завершаю I/O задачу %d: %s\n", task.ID, task.Name)
}

// simulateComputeWork имитирует вычислительную задачу
func simulateComputeWork(task Task, wg *sync.WaitGroup, results chan<- Result) {
	// TODO: Реализуйте функцию для вычислительной работы
	// Подсказки:
	// 1. Не забудьте вызвать wg.Done() в конце функции (используйте defer)
	// 2. Выведите сообщение о начале работы
	// 3. Используйте функцию fibonacci() для создания нагрузки
	// 4. Измерьте время выполнения с помощью time.Now() и time.Since()
	// 5. Создайте объект Result и отправьте его в канал results
	// 6. Выведите сообщение о завершении работы

	defer wg.Done()

	fmt.Printf("🧮 Начинаю вычислительную задачу %d: %s\n", task.ID, task.Name)

	start := time.Now()

	_ = fibonacci(23)

	elapsed := time.Since(start)

	res := Result{
		TaskID:   task.ID,
		TaskName: task.Name,
		Success:  true,
		Duration: elapsed,
		Message:  "Вычисление числа Фибоначчи выполнено"}

	results <- res
}

// fibonacci вычисляет число Фибоначчи (готовая функция)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// monitorProgress показывает прогресс выполнения
func monitorProgress(totalTasks int, results <-chan Result, done chan<- bool) {
	// TODO: Реализуйте функцию мониторинга прогресса
	// Подсказки:
	// 1. Создайте ticker для периодического вывода (каждые 2 секунды)
	// 2. Используйте счетчик для отслеживания завершенных задач
	// 3. Используйте select для обработки ticker и результатов
	// 4. Выводите прогресс в процентах (сколько уже завершилось от общего количества задач)
	// 5. Когда все задачи завершены, отправьте сигнал в канал done

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// TODO: Реализуйте отслеживание количества завершенных задач
	completed := 0

	// Ваш код здесь
	for i := 0; i < totalTasks; i++ {

		
		
		
		go func(id int) {
			
			// Увеличение счетчика завершенных задач
			completed.Add(1)
			fmt.Printf("Задача %d завершена\n", id)\
		}(i)
	}


}

func main() {
	fmt.Println("🚀 === ДЕМОНСТРАЦИЯ ГОРУТИН В GO ===")
	fmt.Println("Запуск параллельных задач с использованием горутин, WaitGroup и каналов")

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// TODO: Создайте списки задач
	ioTasks := []Task{
		{1, "Загрузка данных", "IO"},
		{2, "Чтение файла 1", "IO"},
		{3, "Чтение файла 2", "IO"},
		{4, "Чтение файла 3", "IO"},
		{5, "Чтение файла 4", "IO"},
		{6, "Чтение файла 5", "IO"},
		{7, "Чтение файла 6", "IO"},
		{8, "Чтение файла 7", "IO"},
		{9, "Чтение файла 8", "IO"},
		{10, "Чтение файла 9", "IO"},
		{11, "Чтение файла 10", "IO"},
		{12, "Чтение файла 11", "IO"}}

	computeTasks := []Task{
		{20, "Вычисления", "COMPUTE"},
		{21, "Анализ данных 1", "COMPUTE"},
		{22, "Анализ данных 2", "COMPUTE"},
		{23, "Анализ данных 3", "COMPUTE"},
		{24, "Анализ данных 4", "COMPUTE"},
		{25, "Анализ данных 5", "COMPUTE"},
		{26, "Анализ данных 6", "COMPUTE"},
		{27, "Анализ данных 7", "COMPUTE"},
		{28, "Анализ данных 8", "COMPUTE"},
		{29, "Анализ данных 9", "COMPUTE"}}

	// TODO: Объедините все задачи в один список
	allTasks := append(ioTasks, computeTasks...)
	totalTasks := len(allTasks)

	// TODO: Создайте каналы для результатов и мониторинга
	results := make(chan Result, totalTasks)
	monitorDone := make(chan int)

	// TODO: Создайте WaitGroup для ожидания завершения всех задач
	var wg sync.WaitGroup

	// TODO: Выведите информацию о задачах
	fmt.Printf("\n📋 Запускаю %d задач:\n", totalTasks)
	for _, task := range allTasks {
		fmt.Printf("   • %s (ID: %d, Тип: %s)\n", task.Name, task.ID, task.Type)
	}

	// TODO: Запустите горутину для мониторинга прогресса
	go monitorProgress(totalTasks, results, monitorDone)

	fmt.Println("\n🏃 Запуск горутин...")
	// TODO: Измерьте время выполнения
	startTime := time.Now()

	// TODO: Запустите I/O задачи в горутинах
	// Подсказка: используйте цикл for, wg.Add(1) и go simulateIOWork(...)
	for _, task := range ioTasks {
	    wg.Add(1)
	    go simulateIOWork(task, &wg, results)
	}

	// TODO: Запустите вычислительные задачи в горутинах
	// Подсказка: используйте цикл for, wg.Add(1) и go simulateComputeWork(...)
	for _, task := range computeTasks {
	    wg.Add(1)
	    go simulateComputeWork(task, &wg, results)
	}

	// TODO: Выведите количество запущенных горутин
	fmt.Printf("✨ Запущено %d горутин\n", totalTasks)

	// TODO: Дождитесь завершения всех задач
	wg.Wait()

	// TODO: Закройте канал результатов
	// Подсказка: используйте close(results)

	// TODO: Дождитесь завершения мониторинга
	// Подсказка: читайте из канала monitorDone

	// TODO: Вычислите общее время выполнения
	totalExecutionTime := time.Since(startTime)

	fmt.Printf("\n🎊 === ПРОГРАММА ЗАВЕРШЕНА ===\n")
	// TODO: Выведите статистику выполнения
	fmt.Printf("⏱️  Общее время выполнения программы: %v\n", totalExecutionTime)
	fmt.Printf("🏆 Все горутины успешно завершены!\n")
}
