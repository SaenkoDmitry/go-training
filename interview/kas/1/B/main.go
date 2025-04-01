package main

import (
	"fmt"
	"sort"
)

func Solution(S string) int {
	result := 0
	chOccur := make(map[rune]int)
	for _, ch := range S {
		chOccur[ch]++
	}
	counters := make([]int, 0, len(chOccur))
	for _, count := range chOccur {
		counters = append(counters, count)
	}
	sort.Ints(counters)
	fmt.Println("counters for", S, "=", counters)

	for i := len(counters) - 1; i > 0; i-- {
		if counters[i] == counters[i - 1] {
			counters[i - 1]--
			if counters[i - 1] < 0 {
				continue
			}
			j := i - 1
			if j > 0 && counters[j] < counters[j - 1] {
				i++
				for j > 0 && counters[j] < counters[j - 1] {
					counters[j], counters[j - 1] = counters[j - 1], counters[j]
					j--
				}
			}
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(Solution("aaaabbbb"))
	fmt.Println(Solution("ccaaffddecee"))
	fmt.Println(Solution(""))
	fmt.Println(Solution("a"))
	fmt.Println(Solution("asfklsdf"))
	fmt.Println(Solution("aaaaaaa"))
	fmt.Println(Solution("sdklfsdjfldsklfdsfkljskljsdkflsdjfsdkjldksfjdskl"))
	fmt.Println(Solution("kldsfsj"))
	fmt.Println(Solution("klsdjfs"))
}
