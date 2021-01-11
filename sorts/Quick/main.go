package main

import "fmt"

func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := partition(arr)

	Quicksort(arr[:split])
	Quicksort(arr[split:])
}

func partition(arr []int) int {
	pivot := len(arr) / 2

	left := 0
	right := len(arr) - 1

	for {
		for ; arr[left] < arr[pivot]; left++ {
		}

		for ; arr[right] > arr[pivot]; right-- {
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

func main() {
	arr := []int{3, 2, 1, 5, 4, 0, 15, 13, 10}
	Quicksort(arr)
	fmt.Println(arr)
}
