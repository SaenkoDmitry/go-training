package main

import "fmt"

func zerosToEnd(input []int) []int {
	zeroPointer, nonZeroPointer := 0, 0

	for {
		for zeroPointer < len(input) && input[zeroPointer] != 0 {
			zeroPointer++
		}
		if zeroPointer == len(input) {
			break
		}

		for nonZeroPointer < len(input) && input[nonZeroPointer] == 0 {
			nonZeroPointer++
		}
		if nonZeroPointer == len(input) {
			break
		}

		input[zeroPointer] = input[nonZeroPointer]
		input[nonZeroPointer] = 0
	}
	return input
}

func main() {
	// 0,3,0,1,12 -> 3,1,12,0,0 // нужно передвинуть все нули в конец за O(1) по памяти
	fmt.Println(zerosToEnd([]int{0, 3, 0, 1, 12}))
}
