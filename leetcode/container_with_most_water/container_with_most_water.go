package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	var maxWater, currWater float64
	left := 0
	right := len(height) - 1
	var width int
	for left < right {
		width = right - left
		currWater = float64(width) * math.Min(float64(height[left]), float64(height[right]))
		maxWater = math.Max(maxWater, currWater)
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}
	return int(maxWater)
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
