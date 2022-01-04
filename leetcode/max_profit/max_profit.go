package max_profit

import (
	"math"
)

func MaxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxArr := make([]float64, len(prices))
	minArr := make([]float64, len(prices))

	minArr[0] = float64(prices[0])
	for i := 1; i < len(prices); i++ {
		minArr[i] = math.Min(float64(prices[i]), minArr[i-1])
	}

	maxArr[len(prices) - 1] = float64(prices[len(prices) - 1])
	for i := len(prices) - 2; i >= 1; i-- {
		maxArr[i] = math.Max(float64(prices[i]), maxArr[i+1])
	}

	maxProfit := float64(0)
	for i := 1; i < len(prices); i++ {
		maxProfit = math.Max(maxProfit, maxArr[i] - minArr[i])
	}
	return int(maxProfit)
}
