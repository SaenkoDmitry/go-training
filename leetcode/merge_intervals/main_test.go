package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"contains inner intervals",
			args{
				[][]int{{2, 6}, {2, 3}, {8, 10}, {15, 18}},
			},
			[][]int{{2, 6}, {8, 10}, {15, 18}},
		},
		{
			"reverse order intervals",
			args{
				[][]int{{1, 4}, {0, 4}},
			},
			[][]int{{0, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
