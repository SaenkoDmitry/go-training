package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	worker := func(wg *sync.WaitGroup, mu1, mu2 *sync.Mutex) {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(1 * time.Second)

		mu2.Lock()
		defer mu2.Unlock()
	}

	mu1 := sync.Mutex{}
	mu2 := sync.Mutex{}

	wg.Add(2)
	go worker(&wg, &mu1, &mu2)
	go worker(&wg, &mu2, &mu1)

	wg.Wait()
}
