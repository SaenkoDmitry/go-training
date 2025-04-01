package main

import (
	"reflect"
	"strings"
	"testing"
)

func FillMatrix() [][]string {
	n := 5
	arr := make([][]string, n)
	arr[0] = strings.Split("4 10 7 10 9", " ")
	arr[1] = strings.Split("5 9 0 9 8", " ")
	arr[2] = strings.Split("8 3 6 0 2", " ")
	arr[3] = strings.Split("8 10 3 0 0", " ")
	arr[4] = strings.Split("0 9 0 7 4", " ")
	return arr
}

func TestReverseNumberMinusOne(t *testing.T) {
	arr := FillMatrix()
	result := BypassSpiral(arr)
	expected := strings.Split("6\n0\n9\n0\n0\n3\n10\n3\n9\n10\n7\n10\n9\n8\n2\n0\n4\n7\n0\n9\n0\n8\n8\n5\n4", "\n")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("incorrect results")
	}
}
