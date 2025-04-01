// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// хочется дождаться выполнения функции process
func process(closeCh chan struct{}) chan struct{} {
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)

		for {
			select {
			case <-closeCh:
				return
			default:
				// processing
			}
		}
	}()

	return doneCh
}

func main() {
	closeCh := make(chan struct{})
	doneCh := process(closeCh)

	close(closeCh)
	<-doneCh
	// тут мы точно знаем, что функция закрылась и освободила ресурсы

	fmt.Println("terminated")
}
