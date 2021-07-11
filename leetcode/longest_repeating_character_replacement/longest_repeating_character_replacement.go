package main

import (
	"fmt"
	"math"
)

// AAABBABBAAABBBB


// идея: скользим окном по строке (потому что нам важно найти макс возможную подстроку из повторяющихся символов)
// и если в этом окне разница между наиболее часто встречающимся элементом и длиной окна меньша k, то
// мы понимаешм, что можем обновить максимальную длину
// если же окно становится длиннее, чем максимальн возможная подстрака, состоящая только из повторяющихся символов, то
// это служит поводом подвинуть левую границу
func characterReplacement(s string, k int) int {
	result := 0

	occurrence := make(map[rune]int)
	runes := []rune(s)

	left := 0
	right := 0

	for right < len(runes) {
		occurrence[runes[right]]++

		max, sum := findMaxAndSum(occurrence)
		if sum - max <= k {
			result = int(math.Max(float64(result), float64(right - left + 1)))
		} else {
			occurrence[runes[left]]--
			if val, ok := occurrence[runes[left]]; ok && val == 0 {
				delete(occurrence, runes[left])
			}
			left++
		}

		right++
	}
	return result
}

func findMaxAndSum(m map[rune]int) (max int, sum int) {
	for _, v := range m {
		if v > max {
			max = v
		}
		sum += v
	}
	return max, sum
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}
