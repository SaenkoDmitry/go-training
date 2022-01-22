package main

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example #1",
			args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			20,
		},
		{
			"example #2",
			args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			3,
		},
		{
			"example #3",
			args{
				coins:  []int{2},
				amount: 3,
			},
			-1,
		},
		{
			"example #4",
			args{
				coins:  []int{1},
				amount: 0,
			},
			0,
		},
		{
			"example #5",
			args{
				coins:  []int{4, 7, 8},
				amount: 11,
			},
			2,
		},
		{
			"example #6",
			args{
				coins:  []int{15, 4, 9, 1},
				amount: 18,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
