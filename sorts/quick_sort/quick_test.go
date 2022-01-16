package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			"example #1",
			[]int{33, 70, 10},
			[]int{10, 33, 70},
		},
		{
			"example #2",
			[]int{3, 2, 1, 5, 4, 0, 15, 13, 10},
			[]int{0, 1, 2, 3, 4, 5, 10, 13, 15},
		},
		{
			"example #3",
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			"example #4",
			[]int{1},
			[]int{1},
		},
		{
			"example #5",
			[]int{1,2,2,5},
			[]int{1,2,2,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
