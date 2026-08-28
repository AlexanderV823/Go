package pool

import "time"

// Job представляет задачу для обработки веб-сервера.
type Job struct {
	ID  int
	URL string
}

// Result представляет результат симуляции HTTP-запроса.
type Result struct {
	Job      Job
	Status   string
	Duration time.Duration
}
