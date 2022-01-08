package max_profit_2

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"example #1",
			[]int{7,1,5,3,6,4},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
