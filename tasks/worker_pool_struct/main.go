package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type workerPool struct {
	workersNum int
	ctx        context.Context
	tasks      []func() error
}

type WorkerPoolI interface {
	Go(fn func() error)
	Wait() error
}

func NewWorkerPool(ctx context.Context, workersNum int) WorkerPoolI {
	return &workerPool{
		ctx:        ctx,
		workersNum: workersNum,
		tasks:      make([]func() error, 0),
	}
}

func (w *workerPool) Go(fn func() error) {
	w.tasks = append(w.tasks, fn)
}

func (w *workerPool) Wait() error {
	tasksCh := make(chan func() error, len(w.tasks))
	for i := 0; i < len(w.tasks); i++ {
		tasksCh <- w.tasks[i]
	}
	close(tasksCh)

	errorsCh := make(chan error, len(w.tasks))

	wg := &sync.WaitGroup{}
	for i := 0; i < w.workersNum; i++ {
		wg.Add(1)
		go w.worker(w.ctx, tasksCh, errorsCh, wg)
	}
	wg.Wait()
	close(errorsCh)

	errs := make([]error, 0, len(w.tasks))
	for err := range errorsCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (w *workerPool) worker(ctx context.Context, tasksCh <-chan func() error, errorsCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for fn := range tasksCh {
		select {
		case <-ctx.Done():
			errorsCh <- ctx.Err()
			return
		default:
			if err := fn(); err != nil {
				errorsCh <- err
			}
		}
	}
}

const (
	workersNum = 5
	tasksSize  = 16
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	pool := NewWorkerPool(ctx, workersNum)
	for i := 0; i < tasksSize; i++ {
		i := i
		pool.Go(func() error {
			time.Sleep(time.Second)
			fmt.Printf("%d^2=%d\n", i, i*i)
			if i == 11 {
				return errors.New("11 is magic number")
			}
			return nil
		})
	}

	if err := pool.Wait(); err != nil {
		fmt.Printf("got error: '''%v'''!, elapsed: %s", err, time.Since(start).String())
		return
	}
	fmt.Println("completed successfully. elapsed: ", time.Since(start).String())
}
