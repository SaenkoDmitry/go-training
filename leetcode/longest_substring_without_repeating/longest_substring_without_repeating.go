package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	var maxCount, currCount int
	entries := make(map[rune]struct{})
	var startIndex int
	for i := 0; i < len(runes); i++ {
		if _, exists := entries[runes[i]]; !exists {
			entries[runes[i]] = struct{}{}
			currCount++
			if maxCount < currCount {
				maxCount = currCount
			}
		} else {
			for k := range entries {
				delete(entries, k)
			}
			currCount = 0
			i, startIndex = startIndex, startIndex + 1
		}
	}
	return maxCount
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	s4 := "dvdf"
	s5 := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
	fmt.Println(lengthOfLongestSubstring(s5))
}
