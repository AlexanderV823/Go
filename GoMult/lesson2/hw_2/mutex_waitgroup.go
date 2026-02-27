package main

import (
	"fmt"
	"sync"
	"time"
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

	fmt.Println("Start increment")       // Для визуализации момента запуска
	c.mu.Lock()                          // Установка блокировки
	c.value += increment                 // само увеличение счетчика
	fmt.Printf("Counter: %d\n", c.value) // Визуализация и имитация обработки заблокированных данных
	c.mu.Unlock()                        // снятие блокировки
	time.Sleep(2 * time.Millisecond)     // Если здесь задержка будет меньше, чем в цикле запуска горутин, то вывод в консоль будет последовательным
	fmt.Println("End increment")         // Для визуализации момента завершения
}

func main() {
	// Задание значений
	var numberGorutins int
	var incrementBy int
	fmt.Print("Enter number gorutins: ")
	fmt.Scan(&numberGorutins)
	fmt.Print("Enter incrment: ")
	fmt.Scan(&incrementBy)

	counter := Counter{}
	counter.wg.Add(numberGorutins)

	fmt.Printf("Start %d gorutins...\n", numberGorutins)

	for i := 0; i < numberGorutins; i++ {
		go counter.Increment(incrementBy)
		time.Sleep(1 * time.Millisecond)
	}

	counter.wg.Wait()

	fmt.Printf("End. Final count: %d\n", counter.value)

}
