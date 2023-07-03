package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// MergeIntoIntervals [1,4,5,2,3,9,8,11,0] => "0-5,8-9,11"
//[1,4,3,2] => "1-4"
//[1,4] => "1,4"
func MergeIntoIntervals(arr []int) string {
	sort.Ints(arr)
	groupStart := arr[0]
	groupEnd := arr[0]
	result := make([]string, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1]+1 {
			groupEnd = arr[i]
		} else {
			result = addGroup(result, groupStart, groupEnd)
			groupStart, groupEnd = arr[i], arr[i]
		}
	}
	result = addGroup(result, groupStart, groupEnd)
	return strings.Join(result, ",")
}

func addGroup(result []string, start, end int) []string {
	if start != end {
		result = append(result, fmt.Sprintf("%d-%d", start, end))
	} else {
		result = append(result, strconv.FormatInt(int64(start), 10))
	}
	return result
}

func main() {
	fmt.Println(MergeIntoIntervals([]int{1, 4, 5, 2, 3, 9, 8, 11, 0}))
	fmt.Println(MergeIntoIntervals([]int{1, 4, 3, 2}))
	fmt.Println(MergeIntoIntervals([]int{1, 4}))
}
