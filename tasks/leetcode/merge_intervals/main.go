package main

import (
	"fmt"
	"sort"
)

//type intervalsType [][]int
//
//func (a intervalsType) Len() int {
//	return len([][]int(a))
//}
//
//func (a intervalsType) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//}
//
//func (a intervalsType) Less(i, j []int) bool {
//	if i[0] == j[0] {
//		return i[0] < j[0]
//	}
//	return i[1] < j[1]
//}

func merge(intervals [][]int) [][]int {

	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println(intervals)
	for i := range intervals {

		currentLeft, currentRight := intervals[i][0], intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {

			next := intervals[j]
			if next[0] >= currentLeft && next[0] <= currentRight {
				if next[1] > currentRight {
					currentRight = next[1]
				}
				intervals = append(intervals[:j + 1], intervals[j + 1:]...)
				j--
			} else {
				continue
			}
		}
		result = append(result, []int{currentLeft, currentRight})
	}

	return result
}

func main() {
	intervals := [][]int{{2,6},{1,3},{8,10},{15,18}}
	fmt.Println(merge(intervals))
}
