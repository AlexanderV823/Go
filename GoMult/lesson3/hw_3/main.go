package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID       int
	Customer string
}

// generateOrders - создает список заказов.
// n - количество создаваемых заказов,
// ch - канал со структурой заказа.
func generateOrders(n int, ch chan<- Order) {
	for i := 1; i <= n; i++ {
		order := Order{
			ID:       i,
			Customer: fmt.Sprintf("John Doe №%d", i)}
		ch <- order
	}
	close(ch)
}

// readOrder - читает поступившие заказы из канала
// ch - канал со структурой заказа
// done - канал мониторинга завершения чтения
func readOrder(ch <-chan Order, done chan<- bool) {
	for order := range ch {
		fmt.Printf("Incomming order %d from %s\n", order.ID, order.Customer)
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Incomming order %d validate\n", order.ID)
	}
	done <- true
}

func main() {
	start := time.Now()

	var numberOrders int
	fmt.Print("Enter number generate orders: ")
	fmt.Scan(&numberOrders)

	chOrders := make(chan Order, numberOrders)
	monitorDone := make(chan bool)

	go generateOrders(numberOrders, chOrders)

	go readOrder(chOrders, monitorDone)

	<-monitorDone
	fmt.Println("End.")
	duration := time.Since(start)
	fmt.Printf("Total work time: %s\n", duration)
}
