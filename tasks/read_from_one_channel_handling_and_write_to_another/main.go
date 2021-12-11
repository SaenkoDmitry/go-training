package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for _, num := range []int{1,2,3} {
			naturals <- num
		}
		close(naturals)
	}()

	go func() {
		for id := range naturals {
			squares <- id * id
		}
		close(squares)
	}()

	for id := range squares {
		fmt.Println(id)
	}
}
