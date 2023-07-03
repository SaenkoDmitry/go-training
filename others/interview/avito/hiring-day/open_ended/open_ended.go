package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func do(id int, ch <-chan int, wg *sync.WaitGroup) {
	runtime.Gosched()
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%d done\n", <-ch)
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(1)
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	var wg sync.WaitGroup
	for id := 0; id < 1000; id++ {
		wg.Add(1)
		go do(id, ch, &wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds(), "s")
}
