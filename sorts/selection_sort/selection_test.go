package selection_sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			"example #1",
			[]int{3,2,1},
			[]int{1,2,3},
		},
		{
			"example #2",
			[]int{11,1,4,5, 3, 0, -2, 10},
			[]int{-2, 0, 1, 3, 4, 5, 10, 11},
		},
		{
			"example #3",
			[]int{10},
			[]int{10},
		},
		{
			"example #4",
			[]int{10,10,11,10,10},
			[]int{10,10,10,10,11},
		},
		{
			"example #5",
			[]int{10,101},
			[]int{10,101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args)
			got := tt.args
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
