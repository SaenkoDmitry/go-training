package main

import (
	"fmt"
	"sort"
)

func maxCoins(nums []int) int {
	return maxCoinWithCache(nums, make(map[string]int))
}

func createKey(s []int) string {
	arr := make([]int, 0, len(s))
	arr = append(arr, s...)
	sort.Ints(arr)
	return fmt.Sprintf("%v", arr)
}

func maxCoinWithCache(nums []int, cache map[string]int) int {
	if len(nums) == 0 {
		return 0
	}
	if res, ok := cache[createKey(nums)]; ok {
		return res
	}
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		temp, nums0 := burst(nums[:], i)
		sum := maxCoinWithCache(nums0, cache)
		if temp+sum > maxSum {
			maxSum = temp + sum
		}
	}
	cache[createKey(nums)] = maxSum
	return maxSum
}

func burst(nums []int, i int) (int, []int) {
	if len(nums) == 0 {
		return 1, []int{}
	}
	if len(nums) == 1 {
		return nums[0], []int{}
	}
	if i == 0 {
		return 1 * nums[0] * nums[1], nums[1:]
	}
	if i == len(nums)-1 {
		return 1 * nums[len(nums)-2] * nums[len(nums)-1], nums[:len(nums)-1]
	}
	res := nums[i-1] * nums[i] * nums[i+1]
	nums0 := make([]int, 0, len(nums))
	nums0 = append(nums0, nums[:i]...)
	nums0 = append(nums0, nums[i+1:]...)
	return res, nums0
}

func main() {
	//fmt.Println(maxCoins([]int{3,1,5,8}))
	//fmt.Println(maxCoins([]int{35, 16, 83, 87, 84, 59, 48, 41, 20, 54}))
	fmt.Println(maxCoins([]int{8, 2, 6, 8, 9, 8, 1, 4, 1, 5, 3, 0, 7, 7, 0, 4, 2, 2, 5}))
}
