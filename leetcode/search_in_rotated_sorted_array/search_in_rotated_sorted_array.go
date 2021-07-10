package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := (left + right) / 2

		if nums[pivot] > target {
			right = pivot - 1
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			return pivot
		}
	}
	return -1
}

func search(nums []int, target int) int {

	// normalize with complexity O(N)
	var k int
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i + 1] {
			continue
		}
		k = i + 1
		break
	}
	if k != 0 {
		copy(nums, append(nums[k:], nums[:k + 1]...))
	}

	// binary search with complexity O(logN)
	found := binarySearch(nums, target)
	if found != -1 {
		found = (found + k) % len(nums)
	}
	return found
}

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{3,1}, 3))
}
