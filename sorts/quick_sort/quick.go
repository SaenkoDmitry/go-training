package quick_sort

import (
	"math/rand"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := rand.Intn(len(arr))

	lowPart := make([]int, 0)
	middlePart := make([]int, 0)
	highPart := make([]int, 0)

	for i := range arr {
		if arr[i] < arr[pivot] {
			lowPart = append(lowPart, arr[i])
		} else if arr[i] > arr[pivot] {
			highPart = append(highPart, arr[i])
		} else {
			middlePart = append(middlePart, arr[i])
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart
}
