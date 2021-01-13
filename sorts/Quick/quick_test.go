package main

import (
	"reflect"
	"testing"
)

func handleResult(t *testing.T, actual []int, expected []int) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual array : %v not equal expected : %v", actual, expected)
	}
}

func TestQuicksort1(t *testing.T) {
	given := []int{33, 70, 10}
	expected := []int{10, 33, 70}
	actual := QuickSort(given)

	handleResult(t, actual, expected)
}

func TestQuicksort2(t *testing.T) {
	given := []int{3, 2, 1, 5, 4, 0, 15, 13, 10}
	expected := []int{0, 1, 2, 3, 4, 5, 10, 13, 15}
	actual := QuickSort(given)

	handleResult(t, actual, expected)
}
