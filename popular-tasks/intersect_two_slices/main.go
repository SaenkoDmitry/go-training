package main

import "fmt"

func intersect(a, b []int) (result []int) {
	aSet := make(map[int]int)
	for _, elem := range a {
		aSet[elem]++
	}

	for _, elem := range b {
		if _, ok := aSet[elem]; ok {
			aSet[elem]--
			if val, _ := aSet[elem]; val == 0 {
				delete(aSet, elem)
			}
			result = append(result, elem)
		}
	}
	return
}

func main() {
	a := []int{1,2,3,10,7,8, 8,1}
	b := []int{8,7,1,2,5,4, 1, 1}
	fmt.Println(intersect(a, b))
}
