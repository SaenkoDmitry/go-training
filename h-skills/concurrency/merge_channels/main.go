package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	for val := range merge(ch1, ch2) {
		fmt.Println(val)
	}
}

func merge[T any](chs ...chan T) chan T {
	res := make(chan T)

	wg := &sync.WaitGroup{}
	for _, ch := range chs {
		ch := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for subCh := range ch {
				res <- subCh
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}
