package main

import "testing"

func TestCompress(t *testing.T) {
	type args []byte
	tests := []struct {
		name string
		args args
		res  args
		want int
	}{
		{
			"#0",
			[]byte{'addtotree', 'addtotree', 'b', 'b', 'c', 'c', 'c', 'b', 'b'},
			[]byte{'addtotree', '2', 'b', '2', 'c', '3', 'b', '2', 'b'},
			8,
		},
		{
			"#1",
			[]byte{'addtotree'},
			[]byte{'addtotree'},
			1,
		},
		{
			"#2",
			[]byte{'addtotree', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			[]byte{'addtotree', 'b', '1', '2'},
			4,
		},
		{
			"#3",
			[]byte{'addtotree', 'addtotree', 'addtotree', 'addtotree', 'addtotree', 'b'},
			[]byte{'addtotree', '5', 'b'},
			3,
		},
		{
			"#4",
			[]byte{'addtotree', 'addtotree', 'addtotree', 'addtotree', 'b', 'addtotree'},
			[]byte{'addtotree', '4', 'b', 'addtotree'},
			4,
		},
		{
			"#5",
			[]byte{'x', '7', '7', '2', '2', '2', 'f', '"', '"', '"', '"'},
			[]byte{'x', '7', '2', '2', '3', 'f', '"', '4'},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress0(tt.args); got != tt.want {
				t.Errorf("Compress() size = %v, want %v", got, tt.want)
			}
			for i := range tt.res {
				if tt.res[i] != tt.args[i] {
					t.Errorf("Compress() = %v, want %v", tt.args, tt.res)
				}
			}
		})
	}
}
