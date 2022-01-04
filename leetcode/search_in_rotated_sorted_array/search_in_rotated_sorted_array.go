package search_in_rotated_sorted_array

func SearchInRotatedSortedArrayWithFullScan(nums []int, target int) int {

	// normalize with complexity O(N)
	var k int
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i + 1] {
			continue
		}
		k = i + 1
		break
	}
	if k != 0 {
		copy(nums, append(nums[k:], nums[:k + 1]...))
	}

	// binary search with complexity O(logN)
	found := binarySearch(nums, target)
	if found != -1 {
		found = (found + k) % len(nums)
	}
	return found
}

func SearchInRotatedSortedArray(nums []int, target int) int {
	k := findRotatedIndex(nums)
	nums = append(nums[k:], nums[:k]...)
	if b := binarySearch(nums, target); b == -1 {
		return -1
	} else {
		return (k + b) % len(nums)
	}
}

func findRotatedIndex(nums []int) int {
	left := 0
	right := len(nums) - 1

	if nums[left] <= nums[right] {
		return 0
	}

	last := 0
	for left <= right && nums[left] > nums[right] {
		pivot := (left + right) / 2
		if pivot == left || pivot == right {
			last = pivot + 1
			break
		}
		if nums[pivot] > nums[left] {
			left = pivot + 1
			last = left
		} else if nums[pivot] < nums[left] {
			right = pivot - 1
			last = pivot
		}
	}

	return last
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := (left + right) / 2
		if nums[pivot] < target {
			left = pivot + 1
		} else if nums[pivot] > target {
			right = pivot - 1
		} else {
			return pivot
		}
	}
	return -1
}
