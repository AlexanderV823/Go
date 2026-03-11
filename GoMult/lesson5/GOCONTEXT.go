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
	ID       int
	Duration time.Duration
	Price    int64
	Err      error
	Time     time.Time
}

func processOrder(parentCtx context.Context, wg *sync.WaitGroup, id int, results chan<- OrderResult) {
	defer wg.Done()

	orderCtx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	fmt.Printf("Start processing order %d\n", id)

	// Имитация длительности обработки:
	// часть заказов завершится успешно, часть — по таймауту.
	var workTime time.Duration
	switch id {
	case 2:
		workTime = 4 * time.Second // гарантированно превысит 3 секунды
	case 4:
		workTime = 3500 * time.Millisecond // тоже превысит 3 секунды
	default:
		workTime = time.Duration(500+rand.Intn(2000)) * time.Millisecond
	}

	start := time.Now()

	select {
	case <-time.After(workTime):
		result := OrderResult{
			ID:       id,
			Duration: time.Since(start),
			Price:    rand.Int63n(10000) + 100,
			Err:      nil,
			Time:     time.Now(),
		}
		results <- result
		fmt.Printf("Order %d completed successfully\n", id)

	case <-orderCtx.Done():
		result := OrderResult{
			ID:       id,
			Duration: time.Since(start),
			Price:    0,
			Err:      orderCtx.Err(),
			Time:     time.Now(),
		}
		results <- result
		fmt.Printf("Order %d canceled: %v\n", id, orderCtx.Err())
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	systemCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	go func() {
		<-sigCh
		fmt.Println("Received interrupt signal. Stopping system...")
		cancel()
	}()

	const ordersCount = 5

	results := make(chan OrderResult, ordersCount)
	var wg sync.WaitGroup

	fmt.Println("Order processing system started")

	for i := 1; i <= ordersCount; i++ {
		wg.Add(1)
		go processOrder(systemCtx, &wg, i, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf(
				"Order %-2d | canceled | reason: %-20v | time: %s\n",
				result.ID,
				result.Err,
				result.Time.Format("2006-01-02 15:04:05"),
			)
			continue
		}

		fmt.Printf(
			"Order %-2d | success  | price: %-6d | duration: %-8v | time: %s\n",
			result.ID,
			result.Price,
			result.Duration,
			result.Time.Format("2006-01-02 15:04:05"),
		)
	}

	if systemCtx.Err() != nil {
		fmt.Println("System finished with context status:", systemCtx.Err())
	} else {
		fmt.Println("System finished successfully")
	}
}
