package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"two appearance",
			args{
				"cbaebabacd",
				"abc",
			},
			[]int{0, 6},
		},
		{
			"three appearance",
			args{
				"abab",
				"ab",
			},
			[]int{0, 1, 2},
		},
		{
			"first string less than second",
			args{
				"aaaaaaaaaa",
				"aaaaaaaaaaaaa",
			},
			[]int{},
		},
		{
			"repeat many times only one symbol",
			args{
				"aaaaaaaaaaaaa",
				"aaaaaaaaaa",
			},
			[]int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
