package main

import (
	"fmt"
	"runtime"
	"sync"
)

func joinChannels(channels ...chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		for _, ch := range channels {
			wg.Add(1)
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				runtime.Gosched()
				for id := range ch {
					out <- id
				}
			}(ch, wg)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	numOfChs := 10
	chs := make([]chan int, 0, numOfChs)
	for i := 0; i < numOfChs; i++ {
		chs = append(chs, make(chan int))
	}

	go func() {
		wg := &sync.WaitGroup{}
		for i := range chs {
			wg.Add(1)
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				for j := 0; j < 10; j++ {
					chs[i] <- j + 10*i
				}
				close(chs[i])
			}(i, wg)
		}
		wg.Wait()
	}()

	resultArr := make([]int, 0)
	// ТЗ
	for num := range joinChannels(chs...) {
		resultArr = append(resultArr, num)
		fmt.Println(num)
	}

	grid := make([]int, 100, 100)
	for i := 0; i < len(resultArr); i++ {
		grid[resultArr[i]]++
	}

	fmt.Println("amount:", len(resultArr))
	fmt.Print("check: ")
	for i := 0; i < len(grid); i++ {
		fmt.Print(grid[i])
	}

	// подумать, как сделать, чтобы чтение производилось последовательно из каналов - runtime.Goshed() ?
}
