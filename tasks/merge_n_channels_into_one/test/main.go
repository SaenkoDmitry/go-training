package main

import (
	"fmt"
	"sync"
)

func mergeChs(chs ...chan string) chan string {
	result := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chs {

			wg.Add(1)
			go func() {
				defer wg.Done()

				for c := range ch {
					result <- c
				}
			}()
		}

		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	fmt.Println("Hello, 世界")
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	go func() {
		ch1 <- "11"
		ch1 <- "12"
		ch1 <- "13"
		close(ch1)
	}()
	go func() {
		ch2 <- "21"
		ch2 <- "22"
		close(ch2)
	}()
	go func() {
		ch3 <- "31"
		ch3 <- "32"
		ch3 <- "33"
		close(ch3)
	}()
	go func() {
		ch4 <- "41"
		ch4 <- "42"
		close(ch4)
	}()

	for item := range mergeChs(ch1, ch2, ch3, ch4) {
		fmt.Println(item)
	}
}
