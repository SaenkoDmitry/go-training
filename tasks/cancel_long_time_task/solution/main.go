package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	timeout := 5 * time.Second
	err := executeTaskWithTimeout(timeout)
	if err != nil {
		fmt.Println("elapsed", time.Since(start))
		log.Fatal(err)
	}
}

func executeTaskWithTimeout(t time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	done := make(chan struct{}, 1)

	go func() {
		executeTask()
		done <- struct{}{}
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return errors.New("timeout task")
	}
}

func executeTask() {
	time.Sleep(10 * time.Second)
}
