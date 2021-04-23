package main

import (
	"fmt"
	"sync"
)

func joinChannels(channels ...chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		for _, ch := range channels {
			wg.Add(1)
			go func(ch <- chan int, wg *sync.WaitGroup) {
				for id := range ch {
					out <- id
				}
				wg.Done()
			}(ch, wg)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go func() {
		for _, num := range []int{1,2,3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{10,20,30} {
			b<- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{100,200,300} {
			c<- num
		}
		close(c)
	}()

	// ТЗ
	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}

	// подумать, как сделать, чтобы чтение производилось последовательно из каналов - runtime.Goshed() ?
}
