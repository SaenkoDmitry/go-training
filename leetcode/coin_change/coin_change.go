package main

type Pair struct {
	amount, count int
}

// description here https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	q := make([]*Pair, 0)
	q = append(q, &Pair{amount, 0})

	visited := make(map[int]struct{})
	visited[amount] = struct{}{}

	for len(q) > 0 {
		a, count := q[0].amount, q[0].count
		q = q[1:]

		for i := range coins {
			temp := a - coins[i]
			if temp < 0 {
				continue
			} else if temp == 0 {
				return count + 1
			}

			if _, ok := visited[temp]; !ok {
				q = append(q, &Pair{temp, count + 1})
				visited[temp] = struct{}{}
			}
		}
	}

	return -1
}
