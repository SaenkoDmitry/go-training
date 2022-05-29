package main

import (
	"fmt"
)

// merging N channels into one
func handleChannels(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
	return
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			naturals <- num
		}
		close(naturals)
	}()

	go handleChannels(naturals, squares)

	for i := range squares {
		fmt.Println(i)
	}
}
