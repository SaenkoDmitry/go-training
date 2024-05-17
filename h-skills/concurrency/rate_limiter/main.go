package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// add limiter with limit 10 per second
// read values concurrently
func main() {
	limiter := make(chan struct{}, 10)

	go func() {
		t := time.NewTicker(time.Second)
		for {
			<-t.C
			for i := 0; i < 10; i++ {
				limiter <- struct{}{}
			}
		}
	}()

	count := 50
	ch := make(chan int, count)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- RPCCallWithLimiter(limiter)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println(res)
	}
}

func RPCCall() int {
	return rand.Int()
}

func RPCCallWithLimiter(limiter chan struct{}) int {
	<-limiter
	return RPCCall()
}
