package main

import (
	"fmt"
	//"math/rand/v2" // Можно использовать новый пакет, как рекомендованный для версии начиная с 1.22.
	"math/rand"
	"sync"
	"time"
)

type Metric struct {
	Source string
	Value  float64
	Time   time.Time
}

// cpuMetrics - эмулирует метрики загрузки ЦП
func cpuMetrics() <-chan Metric {

	outCh := make(chan Metric)

	go func() {
		defer close(outCh)

		for i := 0; i < 5; i++ {
			time.Sleep(800 * time.Millisecond)

			// Сразу, без промежуточных переменных, передаем в канал структуру
			outCh <- Metric{
				Source: "CPU",
				Value:  rand.Float64() * 100, // rand.Float64() возвращает число от 0.0 до 1.0. Умножаем на 100, чтобы получить проценты
				Time:   time.Now(),
			}
		}
	}()
	return outCh // Возвращаем созданный канал
}

// memoryMetrics - эмулирует метрики использования памяти
func memoryMetrics() <-chan Metric {

	outCh := make(chan Metric)

	go func() {
		defer close(outCh)

		for i := 0; i < 5; i++ {
			time.Sleep(1200 * time.Millisecond)

			outCh <- Metric{
				Source: "Memory",
				Value:  rand.Float64() * 16384, // Умножаем на 16384, чтобы получить размер в Мб
				Time:   time.Now(),
			}
		}
	}()
	return outCh
}

// networkMetrics - эмулирует метрики сетевой активности
func networkMetrics() <-chan Metric {

	outCh := make(chan Metric)

	go func() {
		defer close(outCh)

		for i := 0; i < 5; i++ {
			time.Sleep(1500 * time.Millisecond)

			outCh <- Metric{
				Source: "Network",
				Value:  rand.Float64() * 1000, // Умножаем на 1000, чтобы получить размер в мегабитах
				Time:   time.Now(),
			}
		}
	}()
	return outCh
}

// fanIn - принимает переменное количество каналов-источников (только для чтения) и создаёт один выходной канал типа chan Metric
func fanIn(channels ...<-chan Metric) <-chan Metric {
	outCh := make(chan Metric)

	var wg sync.WaitGroup

	// Создаем переменную для функции
	mlplx := func(ch <-chan Metric) {
		defer wg.Done()
		for m := range ch {
			outCh <- m
		}
	}

	for _, c := range channels {
		wg.Add(1) // Увеличиваем wg на очередную запущенную функцию
		go mlplx(c)
	}

	go func() {
		wg.Wait()    // Ждем пока все добавленные wg станут Done
		close(outCh) // После того, как все данные в канал переданы (wg.Done), закрываем канал
	}()
	return outCh // Возвращаем канал
}

func main() {

	rand.Seed(time.Now().UnixNano()) // Используем для обратной совместимости со старым пакетом

	fmt.Println("Запуск системы мониторинга…")

	// Создаем общую переменную для всех каналов.
	// В аргументы сразу передаем соответствующие функции, чтобы не создавать промежуточных переменных.
	mergedCh := fanIn(cpuMetrics(), memoryMetrics(), networkMetrics())

	// Выводим в цикле все, что попало в общий канал
	for m := range mergedCh {
		fmt.Printf("Источник: %-8s, Значение: %10.2f, Время: %s\n", m.Source, m.Value, m.Time.Format("2006-01-02 15:04:05.000")) // Используем форматирование и выравнивание
	}

	fmt.Println("Мониторинг завершен.")
}
