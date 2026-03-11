package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type OrderResult struct {
	ID    int
	Price int64
	Time  time.Time
	Err   error
}

// orderProcesed - выполняет имитацию обработки заказа
func orderProcessed(ctxParent context.Context, wg *sync.WaitGroup, id int, results chan<- OrderResult) {
	// Сначало инициализируем defer переданного wg!
	defer wg.Done()

	// Создаем дочерний контекст со своим таймером:
	ctxChild, cancel := context.WithTimeout(ctxParent, 3*time.Second)
	// После создания дочернего контекста инициализируем его defer
	defer cancel()

	fmt.Printf("Start processing order: %d\n", id)

	var workTime time.Duration

	switch id {
	case 1:
		workTime = 3 * time.Second // Для примера заказ №1 будет всегда завершаться по тайм-ауту ctxChild через 3 секунды
	case 2:
		workTime = 10 * time.Second // Для примера заказ №2 будет всегда завершаться по тайм-ауту ctxParent через 10 секунд
	default:
		workTime = 300 * time.Millisecond // Остальные заказы должны успеть обработаться
	}

	select {
	// Отработка получения сигнала о завершении из родительского контекста:
	case <-time.After(workTime):
		results <- OrderResult{
			ID:    id,
			Price: rand.Int63n(10000),
			Time:  time.Now(),
			Err:   nil}
	case <-ctxParent.Done():
		results <- OrderResult{
			ID:    id,
			Price: 0,
			Time:  time.Now(),
			Err:   ctxParent.Err()}
	// Отработка тайм-аута дочернего контекста:
	case <-ctxChild.Done():
		results <- OrderResult{
			ID:    id,
			Price: 0,
			Time:  time.Now(),
			Err:   ctxChild.Err()}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем родительский контекст:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const numberOrders int = 5
	results := make(chan OrderResult, numberOrders) // уфферизированный канал для получения структур с заказами.
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)   // Канал для ожидания сингала отмены выполнения
	signal.Notify(sigCh, os.Interrupt) // Задаем сигнал отмены Ctrl+C, через запятую можно добавить другие сигналы, например, syscall.SIGTERM. Первый параметр канал, куда передавать полученный сигнал.

	go func() {
		<-sigCh // Ждем Ctrl+C
		fmt.Println("Process canceled by user.")
		cancel() // Отменяем контекст для всей программы
	}()

	fmt.Println("Start of order processing...")

	// Цикл для запуска горутин-обработчиков звказов
	for i := 1; i <= numberOrders; i++ {
		wg.Add(1)                               // Перед запуском очередной пополняем счетчик wg
		go orderProcessed(ctx, &wg, i, results) // Передаем родительский контекст, указатель на wg, условный номер заказа и канал куда нужно передать результат обработки
	}

	// Отдельной функцией можно (и нужно!) запустить ожидание завершения всех wg и после закрыть канал в который ожидаются результаты обработки заказов
	go func() {
		wg.Wait()
		close(results)
	}()

	// В цикле читаем канал с результатами:
	for order := range results {
		if order.Err != nil {
			fmt.Printf("Order %-3d: %v at %s\n", order.ID, order.Err, order.Time.Format("2006-01-02 15:04:05.000"))
			continue
		}
		fmt.Printf("Order: %-3d with price: %-10d created at %s\n", order.ID, order.Price, order.Time.Format("2006-01-02 15:04:05.000"))
	}

	// Отрабатываем получение сигнала о завершении из родительского контекста:
	if ctx.Err() != nil {
		fmt.Println("Process:", ctx.Err())
	} else {
		fmt.Println("Done!")
	}
}
