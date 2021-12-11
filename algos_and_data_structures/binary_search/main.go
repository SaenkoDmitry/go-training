package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, searchNumber int) int {
	left, right := 0, len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		guess := arr[mid]

		if guess == searchNumber {
			return mid
		}
		if guess > searchNumber {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 15, 24, 28, 35, 50}
	searchNumber := 24
	sort.Ints(arr)
	founded := binarySearch(arr, searchNumber)
	if founded == -1 {
		fmt.Printf("no such number %d in array %v", searchNumber, arr)
		return
	}
	fmt.Printf("index of number %d is %d in array %v", searchNumber, founded, arr)
}
