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
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'b', 'b'},
			[]byte{'a', '2', 'b', '2', 'c', '3', 'b', '2', 'b'},
			8,
		},
		{
			"#1",
			[]byte{'a'},
			[]byte{'a'},
			1,
		},
		{
			"#2",
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			[]byte{'a', 'b', '1', '2'},
			4,
		},
		{
			"#3",
			[]byte{'a', 'a', 'a', 'a', 'a', 'b'},
			[]byte{'a', '5', 'b'},
			3,
		},
		{
			"#4",
			[]byte{'a', 'a', 'a', 'a', 'b', 'a'},
			[]byte{'a', '4', 'b', 'a'},
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
