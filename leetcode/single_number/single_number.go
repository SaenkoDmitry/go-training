package main

import "fmt"

func singleNumber(nums []int) int {
	mapOfOccurrence := make(map[int]int)

	for i := range nums {
		mapOfOccurrence[nums[i]]++
	}

	for k, v := range mapOfOccurrence {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
	fmt.Println(singleNumber([]int{1}))
}
