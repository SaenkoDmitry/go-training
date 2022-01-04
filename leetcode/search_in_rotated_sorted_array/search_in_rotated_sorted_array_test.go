package search_in_rotated_sorted_array

import "testing"

func Test_findRotatedIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"correct order",
			[]int{1, 2, 3, 4, 5, 6, 7},
			0,
		},
		{
			"center position",
			[]int{5, 1, 3},
			1,
		},
		{
			"minimal right offset near to center",
			[]int{4, 5, 7, 1, 3},
			3,
		},
		{
			"last position",
			[]int{4, 5, 7, 1},
			3,
		},
		{
			"pre-last position",
			[]int{4, 5, 7, 1, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotatedIndex(tt.nums); got != tt.want {
				t.Errorf("findRotatedIndex(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func Test_SearchInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"center position with len=7 with offset=3",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			4,
		},
		{
			"center position with len=7 with offset=3",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			-1,
		},
		{
			"not exist+reversed: center position with len=2 with offset=1",
			args{
				nums:   []int{3, 1},
				target: 3,
			},
			0,
		},
		{
			"not exist: center position with len=2 with offset=1",
			args{
				nums:   []int{1, 3},
				target: 1,
			},
			0,
		},
		{
			"order #0: center position with len=3 with offset=1",
			args{
				nums:   []int{3, 5, 1},
				target: 5,
			},
			1,
		},
		{
			"order #1: center position with len=3 with offset=1",
			args{
				nums:   []int{5, 1, 3},
				target: 1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"find center element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 3,
			},
			2,
		},
		{
			"find first element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 1,
			},
			0,
		},
		{
			"find last element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 5,
			},
			4,
		},
		{
			"find within zero elements",
			args{
				nums:   []int{},
				target: 1,
			},
			-1,
		},
		{
			"find within one element",
			args{
				nums:   []int{1},
				target: 1,
			},
			0,
		},
		{
			"find within one element (not exists)",
			args{
				nums:   []int{3},
				target: 1,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SearchInRotatedSortedArrayWithFullScan(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"find center element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 3,
			},
			2,
		},
		{
			"find first element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 1,
			},
			0,
		},
		{
			"find last element",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 5,
			},
			4,
		},
		{
			"find within zero elements",
			args{
				nums:   []int{},
				target: 1,
			},
			-1,
		},
		{
			"find within one element",
			args{
				nums:   []int{1},
				target: 1,
			},
			0,
		},
		{
			"find within one element (not exists)",
			args{
				nums:   []int{3},
				target: 1,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedSortedArrayWithFullScan(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInRotatedSortedArrayWithFullScan() = %v, want %v", got, tt.want)
			}
		})
	}
}