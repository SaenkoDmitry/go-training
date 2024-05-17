package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// replace batch processing with stream processing
// now program waits until all process functions completes
// it should return jobs concurrently, after processing each job
// use channels
type job struct {
	value int64
	state State
}

type State int

const (
	InitialState State = iota
	FirstStage
	SecondStage
	FinishedStage
)

func FirstProcessing(jobs chan job) chan job {
	result := make(chan job)

	go func() {
		for j := range jobs {
			j.value = int64(float64(j.value) * math.Pi)
			j.state = FirstStage
			result <- j
		}
		close(result)
	}()
	return result
}

func SecondProcessing(jobs chan job) chan job {
	result := make(chan job)

	go func() {
		for j := range jobs {
			j.value = int64(float64(j.value) * math.E)
			j.state = SecondStage
			result <- j
		}
		close(result)
	}()
	return result
}

func LastProcessing(jobs chan job) chan job {
	result := make(chan job)

	go func() {
		for j := range jobs {
			j.value = int64(float64(j.value) / float64(rand.Intn(10)))
			j.state = FinishedStage
			result <- j
		}
		close(result)
	}()
	return result
}

func main() {
	length := 50_000_000
	jobs := make(chan job, length)
	for i := 0; i < length; i++ {
		jobs <- job{int64(i), 0}
	}
	close(jobs)

	start := time.Now()
	jobs = LastProcessing(
		SecondProcessing(
			FirstProcessing(jobs),
		),
	)
	finished := time.Since(start)

	for job := range jobs {
		fmt.Println(job)
	}

	fmt.Println(finished)
}
