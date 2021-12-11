package main

import (
	"fmt"
	"time"
)

type Result struct {
	id  int
	val int
}

func worker(workerID int, in <-chan int, out chan<- *Result) {
	for jobID := range in {
		fmt.Printf("starting execution job=%d in worker=%d\n", jobID, workerID)
		result := jobID * jobID
		time.Sleep(time.Second * 1)
		out <- &Result{
			id: jobID,
			val: result,
		}
		fmt.Printf("finished execution job=%d in worker=%d\n", jobID, workerID)
	}
}

func main() {
	maxWorker := 3
	totalTasks := 10

	in := make(chan int, maxWorker)
	out := make(chan *Result, totalTasks)

	start := time.Now()
	// создание горутин (воркеров)
	for i := 0; i < maxWorker; i++ {
		go worker(i, in, out)
	}

	// добавление задач
	for i := 0; i < totalTasks; i++ {
		in <- i
	}
	close(in)

	for i := 0; i < totalTasks; i++ {
		res := <-out
		fmt.Printf("got result for jobID=%d value=%d\n", res.id, res.val)
	}
	end := time.Now()
	fmt.Printf("elapsed: %f\n", end.Sub(start).Seconds())
}
