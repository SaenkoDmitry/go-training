package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func recursive(nums []int, excluded []int) [][]int {
	if len(nums) == len(excluded) {
		temp := make([]int, 0)
		for _, item := range excluded {
			temp = append(temp, nums[item])
		}
		return [][]int{temp}
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if !contains(excluded, i) {
			res = append(res, recursive(nums, append(excluded, i))...)
		}
	}
	return res
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, recursive(nums, []int{i})...)
	}
	return res
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
}
