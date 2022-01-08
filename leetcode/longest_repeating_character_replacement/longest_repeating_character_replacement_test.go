package main

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example #1",
			args{
				"ABAB",
				2,
			},
			4,
		},
		{
			"example #2",
			args{
				"AABABBA",
				1,
			},
			4,
		},
		{
			"example #3",
			args{
				"ABCBCCBBABABAB",
				1,
			},
			4,
		},
		{
			"example #4",
			args{
				"ABCBCCBBABABAB",
				2,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
