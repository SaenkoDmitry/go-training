package main

import (
	"fmt"
	"sync"
)

func main() {
	var maxElem int = 1000

	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}

	for i := 1000; i > 0; i-- {
		i := i

		wg.Add(1)
		go func() {
			mu.Lock()
			if i%2 == 0 && i < maxElem {
				maxElem = i
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(maxElem)
}
