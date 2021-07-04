package main

import (
	"fmt"
	"sort"
)

type intervalsType [][]int

func (a intervalsType) Len() int {
	return len([][]int(a))
}

func (a intervalsType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a intervalsType) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][0] < a[j][0]
	}
	return a[i][0] < a[j][0]
}

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)

	//sort.Sort(intervalsType(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	for i := 0; i < len(intervals); i++ {
		currentLeft, currentRight := intervals[i][0], intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {

			next := intervals[j]
			if next[0] >= currentLeft && next[0] <= currentRight {
				if next[1] > currentRight {
					currentRight = next[1]
				}
				intervals = append(intervals[:j], intervals[j + 1:]...)
				j--
			} else {
				break
			}
		}
		result = append(result, []int{currentLeft, currentRight})
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
