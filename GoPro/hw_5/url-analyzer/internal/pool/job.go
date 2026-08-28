package pool

import "time"

// Job представляет задачу для обработки URL
type Job struct {
	ID  int
	URL string
}

// Result представляет результат обработки URL
type Result struct {
	JobID    int
	URL      string
	Duration time.Duration
	Status   string
}
