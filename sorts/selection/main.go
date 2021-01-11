package main

import "fmt"

func main() {
	arr := []int{11,1,4,5, 3, 0, -2, 10}
	for i := range arr {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println(arr)
}
