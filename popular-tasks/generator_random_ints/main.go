package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numsGenerator(count int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- r.Intn(count)
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range numsGenerator(10) {
		fmt.Println(num)
	}
}
