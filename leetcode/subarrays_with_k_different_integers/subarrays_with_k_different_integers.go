package main

import "fmt"

// TODO not working properly now
func subarraysWithKDistinct(nums []int, k int) int {
	unique := make(map[int]int)
	var result int

	for left, right := 0, 0; left < len(nums) && right < len(nums); right++ {
		unique[nums[right]]++

		if len(unique) < k {
			continue
		}

		if len(unique) == k {
			result++
			fmt.Println(nums[left:right + 1])
			temp := left

			// двигаем левую границу, пока возможно
			for {
				unique[nums[left]]--
				if unique[nums[left]] == 0 {
					delete(unique, nums[left])
				}
				left++
				if left == len(nums) || left >= right {
					break
				}
				if len(unique) == k {
					result++
					fmt.Println(nums[left:right + 1])
				}
			}
			for left > temp {
				unique[nums[left]]++
				left--
			}
		}

		if len(unique) > k {
			for {
				left++
				if left == len(nums) || left >= right {
					break
				}
				unique[nums[left]]--
				if unique[nums[left]] == 0 {
					delete(unique, nums[left])
				}
				if len(unique) == k {
					result++
					fmt.Println(nums[left:right + 1])
					break
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1,2,1,2,3}, 2))
}
