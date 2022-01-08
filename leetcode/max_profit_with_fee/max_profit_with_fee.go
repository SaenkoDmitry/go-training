package max_profit_with_fee

import "math"

func maxProfit(prices []int, fee int) int {
	buy := make([]int, 0, len(prices))
	sell := make([]int, 0, len(prices))

	buy = append(buy, 0 - prices[0])
	sell = append(sell, 0)

	for i := 1; i < len(prices); i++ {
		prevBuy := buy[i-1]
		currBuy := sell[i-1] - prices[i]
		buy = append(buy, int(math.Max(float64(prevBuy), float64(currBuy))))

		prevSell := sell[i-1]
		currSell := buy[i-1] + prices[i] - fee
		sell = append(sell, int(math.Max(float64(prevSell), float64(currSell))))
	}

	return sell[len(sell)-1]
}
