package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	m := make(map[string][]string)

	for i := range strs {
		tempKey := getSortedStr(strs[i])
		if val, ok := m[tempKey]; ok {
			m[tempKey] = append(val, strs[i])
		} else {
			tempStr := make([]string, 1)
			tempStr[0] = strs[i]
			m[tempKey] = tempStr
		}
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func getSortedStr(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})

	return string(runes)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"addtotree"}))
}
