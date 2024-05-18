package main

import "fmt"

func prefixSum(nums []int, start, end int) int {
	if len(nums) == 0 {
		return 0
	}
	prefix := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		prefix[i] += nums[i] + prefix[i-1]
	}
	return prefix[end] - prefix[start]
}

func main() {
	fmt.Println(prefixSum([]int{1, 2, 3, 4, 5}, 1, 4)) // сумма элементов с индексами 2-4
}
