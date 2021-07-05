package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)
	for i1 := 0; i1 < len(nums); i1++ {

		if i1 != 0 && nums[i1] == nums[i1 - 1] {
			continue
		}
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			if i2 != i1 + 1 && nums[i2] == nums[i2 - 1] {
				continue
			}

			for i3 := i2 + 1; i3 < len(nums); i3++ {
				if i3 != i2 + 1 && nums[i3] == nums[i3 - 1] {
					continue
				}

				for i4 := i3 + 1; i4 < len(nums); i4++ {
					if i4 != i3 + 1 && nums[i4] == nums[i4 - 1] {
						continue
					}

					if nums[i1] + nums[i2] + nums[i3] + nums[i4] == target {
						result = append(result, []int{nums[i1], nums[i2], nums[i3], nums[i4]})
					}
				}
			}
		}
	}

	return result
}


func main() {
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
	fmt.Println(fourSum([]int{2,2,2,2,2}, 8))
}
