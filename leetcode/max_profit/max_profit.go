package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	minArr := make([]int, len(prices))
	maxArr := make([]int, len(prices))

	min := prices[0]
	minArr[0] = min
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		minArr[i] = min
	}

	max := prices[len(prices) - 1]
	maxArr[0] = max
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		maxArr[i] = max
	}

	var maxDiff, currDiff int
	for i := 0; i < len(prices); i++ {
		currDiff = maxArr[i] - minArr[i]
		if currDiff > maxDiff {
			maxDiff = currDiff
		}
	}
	return maxDiff
}

func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
	prices = []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
	prices = []int{1,2}
	fmt.Println(maxProfit(prices))
}
