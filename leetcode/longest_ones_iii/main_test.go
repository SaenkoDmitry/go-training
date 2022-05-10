package longest_ones_iii

import "testing"

func TestLongestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"#0",
			args{
				[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				2,
			},
			6,
		},
		{
			"#1",
			args{
				[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				3,
			},
			10,
		},
		{
			"#2",
			args{
				[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				10,
			},
			19,
		},
		{
			"#2",
			args{
				[]int{1, 1, 0},
				1,
			},
			3,
		},
		{
			"#3",
			args{
				[]int{0, 0, 0, 0},
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("LongestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
