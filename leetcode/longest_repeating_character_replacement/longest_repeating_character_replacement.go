package main

import (
	"fmt"
	"math"
)

// AAABBABBAAABBBB


// идея: скользим окном по строке (потому что нам важно найти макс возможную подстроку из повторяющихся символов)
// и если в этом окне разница между наиболее часто встречающимся элементом и длиной окна меньша k, то
// мы понимаем, что можем обновить максимальную длину (подвинуть правую границу)
// если же окно становится длиннее, чем максимальн возможная подстрока, состоящая только из повторяющихся символов, то
// это служит поводом подвинуть левую границу
func characterReplacement(s string, k int) int {
	runes := []rune(s)

	result := 0
	left, right := 0, 0
	m := make(map[rune]int)
	max := 0
	for right < len(runes) {
		m[runes[right]]++
		if val := m[runes[right]]; val > max {
			max = val
		}
		if max + k < right - left + 1 {
			m[runes[left]]--
			if val, ok := m[runes[left]]; ok && val == 0 {
				delete(m, runes[left])
			}
			left++
		} else {
			result = int(math.Max(float64(result), float64(right - left + 1)))
		}

		right++
	}
	return result
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}
