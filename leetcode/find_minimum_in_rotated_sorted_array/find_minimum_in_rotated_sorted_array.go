package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	n := findRotatedIndex(nums)
	return nums[n]
}

func findRotatedIndex(nums []int) int {
	left := 0
	right := len(nums) - 1

	result := 0
	for left < right && nums[left] > nums[right] {
		pivot := (left + right) / 2

		if nums[left] > nums[pivot] {
			right = pivot - 1
			result = pivot
		} else if nums[left] <= nums[pivot] {
			left = pivot + 1
			result = pivot + 1
		}
	}
	return result
}
