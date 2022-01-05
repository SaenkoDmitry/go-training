package find_minimum_in_rotated_sorted_array

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"center element #1",
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			"center element #2",
			[]int{4,5,6,7,0,1,2},
			0,
		},
		{
			"start of array",
			[]int{11,13,15,17},
			11,
		},
		{
			"two elements",
			[]int{11,10},
			10,
		},
		{
			"one element",
			[]int{11},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
