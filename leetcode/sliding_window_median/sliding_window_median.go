package main

import (
	"fmt"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	sortedArr := make([]int, k)
	copy(sortedArr, nums[:k])
	sort.Ints(sortedArr)

	result := make([]float64, 0, len(nums) - k + 1)

	result = append(result, getMedian(sortedArr))
	for i := k; i < len(nums); i++ {
		lastIndex := i - k
		var replaceIndex int
		for replaceIndex = 0; replaceIndex < len(sortedArr); replaceIndex++ {
			if sortedArr[replaceIndex] == nums[lastIndex] {
				break
			}
		}

		sortedArr[replaceIndex] = nums[i]
		for ; replaceIndex > 0 && sortedArr[replaceIndex] < sortedArr[replaceIndex - 1]; replaceIndex-- {
			sortedArr[replaceIndex], sortedArr[replaceIndex - 1] = sortedArr[replaceIndex - 1], sortedArr[replaceIndex]
		}
		for ; replaceIndex < len(sortedArr) - 1 && sortedArr[replaceIndex] > sortedArr[replaceIndex + 1]; replaceIndex++ {
			sortedArr[replaceIndex], sortedArr[replaceIndex + 1] = sortedArr[replaceIndex + 1], sortedArr[replaceIndex]
		}
		result = append(result, getMedian(sortedArr))
	}
	return result
}

func getMedian(nums []int) float64 {
	if len(nums) % 2 == 0 {
		return float64(nums[len(nums) / 2 - 1] + nums[len(nums) / 2]) / float64(2)
	} else {
		return float64(nums[len(nums) / 2])
	}
}

func main() {
	fmt.Println(medianSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(medianSlidingWindow([]int{1,4,2,3}, 4))
}
