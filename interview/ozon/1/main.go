package main

import "fmt"

func exists(arr1 []int, arr2 []int) []int {
	m := make(map[int]struct{})
	for i := range arr1 {
		m[arr1[i]] = struct{}{}
	}

	result := make([]int, 0)
	for i := range arr2 {
		if _, ok := m[arr2[i]]; ok {
			result = append(result, arr2[i])
		}
	}
	return result
}

func main() {
	fmt.Println(exists([]int{1,2,3,4,5}, []int{4,5}))
	fmt.Println(exists([]int{4,5}, []int{4,5}))
	fmt.Println(exists([]int{4}, []int{4,5}))
}
