package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	leackyBucketCh chan struct{}

	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewLeakyBucketLimiter(limit int, period time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		leackyBucketCh: make(chan struct{}, limit),
		closeCh:        make(chan struct{}),
		closeDoneCh:    make(chan struct{}),
	}

	leakInterval := period.Nanoseconds() / int64(limit)
	go limiter.StartPeriodicLeak(time.Duration(leakInterval))
	return limiter
}

func (l *RateLimiter) StartPeriodicLeak(interval time.Duration) {
	timer := time.NewTimer(interval)
	defer func() {
		timer.Stop()
		close(l.closeDoneCh)
	}()

	for {
		select {
		case <-l.closeCh:
			return
		default:
		}

		select {
		case <-l.closeCh:
			return
		case <-timer.C:
			select {
			case <-l.leackyBucketCh:
			default:
			}
		}
	}
}

func (l *RateLimiter) Allow() bool {
	select {
	case l.leackyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l *RateLimiter) Shutdown() {
	close(l.closeCh)
	<-l.closeDoneCh
}

func main() {
	limiter := NewLeakyBucketLimiter(5, time.Second)
	for i := 0; i < 6; i++ {
		time.Sleep(200 * time.Millisecond)
		limiter.Allow()
		fmt.Println("i=", i)
	}
}
