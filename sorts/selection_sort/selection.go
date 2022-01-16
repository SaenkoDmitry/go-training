package selection_sort

func SelectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
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
}
