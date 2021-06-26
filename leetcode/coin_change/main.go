package main

import "fmt"

func coinChangeRecursive(accum int, coins []int, amount int) int {
	min := 0
	found := false

	if amount == 0 {
		return accum
	}

	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			temp := coinChangeRecursive(accum + 1, coins, amount - coins[i])
			if !found || temp < min {
				found = true
				min = temp
			}
		}
	}

	if !found {
		return -1
	}

	return min
}

func coinChange(coins []int, amount int) int {
	return coinChangeRecursive(0, coins, amount)
}


func main() {
	fmt.Println(coinChange([]int{1,2,5}, 11))
}
