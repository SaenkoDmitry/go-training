package main

import "fmt"

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		arr = merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
	}
	return arr
}

func main() {
	arr := []int{3, 2, 1, 5, 4, 0, 15, 13, 10}
	arr = mergeSort(arr)
	fmt.Println(arr)
}
