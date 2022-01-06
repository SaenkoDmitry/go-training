package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	for i := range intervals {
		left, right := intervals[i][0], intervals[i][1]
		found := false
		for j := 0; j < len(result); j++ {
			if result[j][0] <= left && left <= result[j][1] {
				if result[j][1] < right {
					result[j][1] = right
				}
				found = true
			}
		}
		if !found {
			result = append(result, []int{left, right})
		}
	}
	return result
}


func main() {
	//intervals := [][]int{{2,6},{1,3},{8,10},{15,18}}
	intervals2 := [][]int{{2,6},{2,3},{8,10},{15,18}}
	//intervals := [][]int{{1,4},{4,5}}
	//intervals := [][]int{{1,4},{0,4}}
	//intervals := [][]int{{1,4},{2,3}}
	//fmt.Println(merge(intervals))
	fmt.Println(merge(intervals2))
}
