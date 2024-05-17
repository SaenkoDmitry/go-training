package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is from main function!")

	quitChan := make(chan bool)
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				// print a message every 3 seconds
				fmt.Println("This is test goroutine")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	// sleep to print some message from the goroutine
	time.Sleep(time.Second * 10)

	// stop the goroutine
	quitChan <- true
	fmt.Println("Go rountine has stopped!")

	// test if the gorountine stopped or not
	time.Sleep(time.Second * 10)
	fmt.Println("This is the end of main func, goroutine no longer run!")
}
