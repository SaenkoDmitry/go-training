package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			panic(err)
		}
	}

	curr := 0
	var best float64
	for _, elem := range arr {
		if elem == 1 {
			curr++
			best = math.Max(float64(curr), best)
		} else {
			curr = 0
		}
	}

	fmt.Println(best)
}
