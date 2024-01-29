package main

import "fmt"

func removeZeros(nums []int) []int {
	nums0 := nums[:0]
	for i := range nums {
		if nums[i] != 0 {
			nums0 = append(nums0, nums[i])
		}
	}
	return nums0
}

func main() {
	fmt.Println(removeZeros([]int{1, 0, 7, 0, 0, 42}))
}
