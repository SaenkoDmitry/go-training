package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var addresses = []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}

func get(ctx context.Context, address, key string) (string, error) { // already implemented
	// time.Sleep(time.Second * 10) // first option
	// return "", errNotFound       // second option
	// return "", errors.New("failed") // third option
	return "value", nil // fourth option
}

var errNotFound = errors.New("not found")
var distributedErr = errors.New("something went wrong")

func distributedGet(ctx context.Context, addrs []string, key string) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	valueCh := make(chan string)
	notFoundErrCh := make(chan error)

	wg := &sync.WaitGroup{}

	for i := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()

			var v string
			var err error

			withRetry(func() error {
				v, err = get(ctx, addr, key)
				return err
			}, 3, time.Millisecond*10)

			if errors.Is(err, errNotFound) {

				select {
				case notFoundErrCh <- err:
				default:
				}

			} else if err != nil {
				return
			}

			select {
			case valueCh <- v:
			default:
			}

		}(addrs[i])
	}

	go func() {
		wg.Wait()
		close(valueCh)
		close(notFoundErrCh)
	}()

	select {
	case v, ok := <-valueCh:
		if !ok {
			return "", distributedErr
		}
		return v, nil
	case err, ok := <-notFoundErrCh:
		if !ok {
			return "", distributedErr
		}
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func withRetry(cb func() error, count int, delay time.Duration) {
	err := cb()

	for i := 0; i < count; i++ {
		if err != nil && !errors.Is(err, errNotFound) {
			err = cb()
		}
		time.Sleep(delay)
		delay = delay * 2
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := distributedGet(ctx, addresses, "key")
	fmt.Println("res:", res, "error:", err)
}
