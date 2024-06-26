package longest_ones_iii

// LongestOnes идея: два указателя, двигаем сначала правую границу, затем если больше не можем двигать, то двигаем левую
func LongestOnes(nums []int, k int) int {
	maxSeq := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] != 1 {
			if k > 0 {
				k--
			} else {
				for left < len(nums) && nums[left] > 0 {
					left++
				}
				left++
			}
		}
		if right-left+1 > maxSeq {
			maxSeq = right - left + 1
		}
	}
	return maxSeq
}
