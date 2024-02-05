package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Result struct {
	id  int
	val int
}

func worker(ctx context.Context, wg *sync.WaitGroup, workerID int, in <-chan int, out chan<- *Result) {
	defer wg.Done()

	for jobID := range in {
		fmt.Printf("starting execution job=%d in worker=%d\n", jobID, workerID)
		result := jobID * jobID
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			return
		case out <- &Result{
			id:  jobID,
			val: result,
		}:
			fmt.Println(jobID, "executing....")
		}
		//fmt.Printf("finished execution job=%d in worker=%d\n", jobID, workerID)
	}
}

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	maxWorker := 3
	totalTasks := 20

	in := make(chan int, totalTasks)
	out := make(chan *Result, totalTasks)

	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU() - 1))

	wg := &sync.WaitGroup{}

	start := time.Now()
	// создание горутин (воркеров)
	for i := 0; i < maxWorker; i++ {
		wg.Add(1)
		go worker(ctx, wg, i, in, out)
	}

	// добавление задач
	for i := 0; i < totalTasks; i++ {
		in <- i
	}
	close(in)

	go func() {
		wg.Wait()
		close(out)
	}()

	fmt.Println("=====")
	fmt.Println("collecting results:")
	count := 0
	for job := range out {
		fmt.Printf("got result for jobID=%d value=%d\n", job.id, job.val)
		count++
	}
	fmt.Println("=====")
	fmt.Println("completed", count, "jobs of", totalTasks, "in", timeout)
	end := time.Now()
	fmt.Printf("elapsed: %f\n", end.Sub(start).Seconds())
}
