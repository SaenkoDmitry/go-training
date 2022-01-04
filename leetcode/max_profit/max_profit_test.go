package max_profit

import "testing"

func Test_MaxProfit(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"diff between two center elements",
			[]int{7,1,5,3,6,4},
			5,
		},
		{
			"descending order",
			[]int{7,6,4,3,1},
			0,
		},
		{
			"two elements",
			[]int{1,2},
			1,
		},
		{
			"diff between center element and last",
			[]int{6,1,3,2,4,7},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args); got != tt.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
