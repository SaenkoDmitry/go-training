// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func orDone(inputCh chan string, doneCh chan struct{}) chan string {
	outputCh := make(chan string)

	go func() {
		defer close(outputCh)

		for {
			select {
			case <-doneCh:
				return
			default:
			}

			select {
			case value, opened := <-inputCh:
				if !opened {
					return
				}
				outputCh <- value
			case <-doneCh:
				return
			}
		}
	}()

	return outputCh
}

func main() {
	channel := make(chan string)

	go func() {
		for {
			channel <- "test"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		close(done)
	}()

	for value := range orDone(channel, done) {
		fmt.Println(value)
	}

	/*
		for {
			select {
			case val, opened := <- inputCh:
				if !opened {
					return
				}
				// processing ...
			case <- doneCh:
				return
		}
	*/
}
