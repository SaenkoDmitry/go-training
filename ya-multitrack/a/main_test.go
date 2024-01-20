package main

import "testing"

func Test_maxDigitsCount(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: 15,
				b: 2000,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDigitsCount(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maxDigitsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "12345",
			},
			want: "54321",
		},
		{
			args: args{
				s: "1",
			},
			want: "1",
		},
		{
			args: args{
				s: "",
			},
			want: "",
		},
		{
			args: args{
				s: "10000",
			},
			want: "00001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseNumber(t *testing.T) {
	type args struct {
		input       int
		digitsCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				input:       105,
				digitsCount: 4,
			},
			want: 5010,
		},
		{
			args: args{
				input:       100,
				digitsCount: 5,
			},
			want: 100,
		},
		{
			args: args{
				input:       100,
				digitsCount: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNumber(tt.args.input, tt.args.digitsCount); got != tt.want {
				t.Errorf("reverseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
