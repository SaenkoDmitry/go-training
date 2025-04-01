package main

import "fmt"

func generator(start, end, step int) <-chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)
		for i := start; i < end; i += step {
			outputCh <- i
		}
	}()

	return outputCh
}

func main() {
	for number := range generator(100, 200, 5) {
		fmt.Println(number)
	}
}
