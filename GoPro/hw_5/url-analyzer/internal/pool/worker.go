package pool

import (
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		delay := time.Duration(100+rand.Intn(501)) * time.Millisecond
		time.Sleep(delay)

		// Инициализация структуры Result
		results <- Result{
			JobID:    job.ID, // Присваиваем ID задачи в поле JobID
			URL:      job.URL,
			Duration: delay,
			Status:   "Успешно",
		}
	}
}

type Pool struct {
	workerCount int
}

func NewPool(workerCount int) *Pool {
	if workerCount < 1 {
		workerCount = 1
	}
	return &Pool{workerCount: workerCount}
}

func (p *Pool) Start(jobs []Job) []Result {
	if len(jobs) == 0 {
		return nil
	}

	jobsChan := make(chan Job, len(jobs))
	resultsChan := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	for w := 1; w <= p.workerCount; w++ {
		wg.Add(1)
		go worker(w, jobsChan, resultsChan, &wg)
	}

	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	wg.Wait()
	close(resultsChan)

	var results []Result
	for res := range resultsChan {
		results = append(results, res)
	}

	return results
}
