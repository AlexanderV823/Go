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
func orderProcesed(ctxParent context.Context, wg *sync.WaitGroup, id int, results chan<- OrderResult) {
	defer wg.Done()

	ctxChild, cancel := context.WithTimeout(ctxParent, 3*time.Second)
	defer cancel()

	fmt.Printf("Start processing order: %d\n", id)
	for i := 0; i < 3; i++ {
		select {
		case <-ctxParent.Done():
			results <- OrderResult{
				ID:    id,
				Price: 0,
				Time:  time.Now(),
				Err:   ctxParent.Err()}
			return
		case <-ctxChild.Done():
			results <- OrderResult{
				ID:    id,
				Price: 0,
				Time:  time.Now(),
				Err:   ctxChild.Err()}
			return
		default:
			switch {
			case id == 1:
				time.Sleep(3 * time.Second) // Для примера заказ №1 будет всегда завершаться по таймауту ctxChild через 3 секунды
			case id == 2:
				time.Sleep(10 * time.Second) // Для примера заказ №2 будет всегда завершаться по таймауту ctxParent через 10 секунд
			default:
				time.Sleep(300 * time.Millisecond) // Остальные заказы должны успеть обработаться
			}
		}
	}

	results <- OrderResult{
		ID:    id,
		Price: rand.Int63n(10000),
		Time:  time.Now(),
		Err:   nil}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const numberOrders int = 5
	results := make(chan OrderResult, numberOrders)
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt) // Задаем сигнал отмены Ctrl+C

	go func() {
		<-sigCh // Ждем Ctrl+C
		fmt.Println("Process canceled by user.")
		cancel() // Отменяем контекст для всей программы
	}()

	fmt.Println("Start of order processing...")

	for i := 1; i <= numberOrders; i++ {
		wg.Add(1)
		go orderProcesed(ctx, &wg, i, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for order := range results {
		if order.Err != nil {
			fmt.Printf("Order %-3d: %v at %s\n", order.ID, order.Err, order.Time.Format("2006-01-02 15:04:05.000"))
			continue
		}
		fmt.Printf("Order: %-3d with price: %-10d created at %s\n", order.ID, order.Price, order.Time.Format("2006-01-02 15:04:05.000"))
	}

	if ctx.Err() != nil {
		fmt.Println("Process:", ctx.Err())
	}
}
