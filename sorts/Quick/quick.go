package main

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := rand.Intn(len(arr))

	lowPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))

	median := arr[pivot]

	for _, elem := range arr {
		switch {
		case elem < median:
			lowPart = append(lowPart, elem)
		case elem > median:
			highPart = append(highPart, elem)
		case elem == median:
			middlePart = append(middlePart, elem)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

// not works
func Quicksort2(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := partition(arr)

	Quicksort2(arr[:split])
	Quicksort2(arr[split:])
}

func partition(arr []int) int {
	pivot := len(arr) / 2

	left := 0
	right := len(arr) - 1

	for left <= right {
		for ; arr[left] < arr[pivot]; left++ { }
		for ; arr[right] > arr[pivot]; right-- { }

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
		right--
		left++
	}
	return right
}
