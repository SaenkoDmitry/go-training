package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
	fmt.Println(search([]int{5}, 5))
}
