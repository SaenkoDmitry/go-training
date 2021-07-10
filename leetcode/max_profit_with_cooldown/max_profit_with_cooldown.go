package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// максимум, пред-предыдущем шаге
	bestPrevPrev := float64(0)

	// максимум на предыдущем шаге
	bestPrev := math.Max(0, float64(prices[1]-prices[0]))

	// запоминаем текущее значение прибыли, если купили минимальный элемент
	tempWhenBuyMin := math.Max(float64(-prices[0]), float64(-prices[1]))

	for i := 2; i < len(prices); i++ {
		// либо оставляем прибыль с прошлого шага, либо фиксируем сделку
		best := math.Max(bestPrev, tempWhenBuyMin+float64(prices[i]))

		// решаем, что выгоднее: остаться в состоянии покупку некоторого элемента, либо зафиксировать cooldown
		tempWhenBuyMin = math.Max(tempWhenBuyMin, bestPrevPrev-float64(prices[i]))

		// обновляем предыдущие максимумы
		bestPrevPrev = bestPrev
		bestPrev = best
	}
	return int(bestPrev)
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
