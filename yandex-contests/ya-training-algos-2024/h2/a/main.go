package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pair struct {
	left  int
	right int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int

	fmt.Fscan(in, &n)

	pairs := make([]Pair, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		pairs = append(pairs, Pair{a, b})
	}

	minFirst, minSecond := math.MaxInt, math.MaxInt
	maxFirst, maxSecond := 0, 0
	for _, pair := range pairs {
		if pair.left < minFirst {
			minFirst = pair.left
		}
		if pair.left > maxFirst {
			maxFirst = pair.left
		}
		if pair.right < minSecond {
			minSecond = pair.right
		}
		if pair.right > maxSecond {
			maxSecond = pair.right
		}
	}
	fmt.Println(minFirst, minSecond, maxFirst, maxSecond)
}
