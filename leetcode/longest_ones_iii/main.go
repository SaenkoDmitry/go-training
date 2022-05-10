package longest_ones_iii

// LongestOnes идея: два указателя, двигаем сначала правую границу, затем если больше не можем двигать, то двигаем левую
func LongestOnes(nums []int, k int) int {
	left, count, max := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 1 {
			if count < k {
				count++
			} else {
				for {
					if nums[left] == 0 {
						left++
						break
					}
					left++
				}
			}
		}
		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max
}
