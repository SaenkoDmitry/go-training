package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			// counter = counter + 1 // it's race
			// to see problem you can run command : 'go run -race tasks/race_condition/main.go'
			atomic.AddInt64(&counter, 1)
			/*
				// another way to sync add operation below
				for !atomic.CompareAndSwapInt64(&counter, counter, counter+1) {
				}
			*/
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
