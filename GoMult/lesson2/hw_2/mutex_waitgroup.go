package main

import (
	"fmt"
	"sync"
)

// Объединение счетчик, мьютекса и wg в одну структуру - хорошая практика
type Counter struct {
	mu    sync.Mutex
	wg    sync.WaitGroup
	value int
}

// Увеличивает значение счетчика на 1000
func (c *Counter) Increment(increment int) {
	defer c.wg.Done()

	fmt.Println("Start increment") // Для визуализации момента запуска

	for j := 0; j < increment; j++ {
		c.mu.Lock()   // Установка блокировки
		c.value++     // само увеличение счетчика
		c.mu.Unlock() // снятие блокировки
	}

	fmt.Println("End increment") // Для визуализации момента завершения
}

func main() {
	// Задание значений
	var numberGorutins int
	var incrementBy int
	fmt.Print("Enter number gorutins: ")
	fmt.Scan(&numberGorutins)
	// fmt.Print("Enter incrment: ")
	// fmt.Scan(&incrementBy)
	incrementBy = 1000

	counter := Counter{}
	counter.wg.Add(numberGorutins)

	fmt.Printf("Start %d gorutins...\n", numberGorutins)

	for i := 0; i < numberGorutins; i++ {
		go counter.Increment(incrementBy)
	}

	counter.wg.Wait()

	fmt.Printf("End. Final count: %d\n", counter.value)
}
