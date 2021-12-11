package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, f func(int, int) int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	fmt.Printf("worker with id = %d started\n", id)
	for j := range jobs {
		runtime.Gosched()
		calc := f(j, 10)
		result <- calc
		fmt.Printf("worker with id = %d calculated result = %d\n", id, calc)
	}
	fmt.Printf("worker with id = %d finished\n", id)
	wg.Done()
}

func main() {
	var startTime, endTime time.Time

	startTime = time.Now()
	const numOfJobs = 100
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	multiplier := func(x, y int) int {
		return x * y
	}

	// запускаем воркеры
	const numOfWorkers = 3
	wg := &sync.WaitGroup{}
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, multiplier, jobs, results, wg)
	}

	// кладем данные
	for i := 0; i < numOfJobs; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// достаем обработанные результаты
	for i := 0; i < numOfJobs; i++ {
		fmt.Println(<-results)
	}
	close(results)

	endTime = time.Now()

	fmt.Printf("work take %d microseconds", endTime.Sub(startTime).Microseconds())
}
