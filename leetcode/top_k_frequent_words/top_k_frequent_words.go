package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	occurrences := make(map[string]int)
	for i := range words {
		occurrences[words[i]]++
	}

	keys := make([]string, 0, len(occurrences))
	for k := range occurrences {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		firstWord := keys[i]
		secondWord := keys[j]
		firstCount, _ := occurrences[firstWord]
		secondCount, _ := occurrences[secondWord]
		if firstCount == secondCount {
			return firstWord < secondWord
		}
		return firstCount > secondCount
	})

	result := make([]string, 0)

	for i := 0; i < k; i++ {
		result = append(result, keys[i])
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
