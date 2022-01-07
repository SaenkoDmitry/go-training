package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example #1",
			args{
				[]int{1, 1, 1, 2, 2, 3},
				2,
			},
			[]int{1, 2},
		},
		{
			"example #2",
			args{
				[]int{4,1,-1,2,-1,2,3},
				2,
			},
			[]int{-1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
