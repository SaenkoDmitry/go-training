package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}
	sort.Ints(A)

	if A[len(A) - 1] <= 0 {
		return 1
	}

	prev := A[0]
	for _, item := range A {
		if item != prev + 1 && item != prev {
			return prev + 1
		}
		prev = item
	}
	return A[len(A) - 1] + 1
}

func main() {
	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2}))
	fmt.Println(Solution([]int{-1, -3}))
}
