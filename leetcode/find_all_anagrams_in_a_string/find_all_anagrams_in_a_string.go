package find_all_anagrams_in_a_string

import "reflect"

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)

	runes := []rune(p)
	k := len(runes)
	m := make(map[rune]int)
	for i := range runes {
		m[runes[i]]++
	}

	window := make(map[rune]int)
	runes = []rune(s)
	for i := 0; i < len(runes); i++ {
		window[runes[i]]++
		if i >= k {
			window[runes[i-k]]--
			if val, ok := window[runes[i-k]]; ok && val == 0 {
				delete(window, runes[i-k])
			}
		}
		if i >= k - 1 && reflect.DeepEqual(window, m) {
			result = append(result, i-k+1)
		}
	}
	return result
}

