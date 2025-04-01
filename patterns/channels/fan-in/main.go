// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func mergeChs[T any](chs ...chan T) chan T {
	resultCh := make(chan T)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))

		for _, ch := range chs {
			go func() {
				defer wg.Done()
				for value := range ch {
					resultCh <- value
				}
			}()
		}

		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i += 3 {
			ch1 <- i
			ch2 <- i + 1
			ch3 <- i + 2
		}
	}()

	for value := range mergeChs(ch1, ch2, ch3) {
		fmt.Println("value:", value)
	}
}
