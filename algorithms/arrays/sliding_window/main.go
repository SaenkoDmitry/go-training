package main

import "fmt"

func longestConsecutiveOnes(nums []int, k int) int {
	maxSeq := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] != 1 {
			if k > 0 {
				k--
			} else {
				for left < len(nums) && nums[left] > 0 {
					left++
				}
				left++
			}
		}
		if right-left+1 > maxSeq {
			maxSeq = right - left + 1
		}
	}
	return maxSeq
}

func main() {
	fmt.Println(4 == longestConsecutiveOnes([]int{1, 0, 0, 1, 0, 0, 0}, 2))
}
