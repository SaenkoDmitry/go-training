package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if reflect.DeepEqual(a, b) {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestGetTopWhenIsEmpty(t *testing.T) {
	binaryHeap := NewBinaryHeap()
	if _, err := binaryHeap.Top(); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Fatal("errors not produced")
	}
}

func TestGetTopWhenFull(t *testing.T) {
	binaryHeap := NewBinaryHeap()
	binaryHeap.Push(1)
	binaryHeap.Push(4)
	binaryHeap.Push(3)
	binaryHeap.Push(5)
	binaryHeap.Push(15)
	var topElement int
	var err error
	if topElement, err = binaryHeap.Top(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(topElement)
	}
	assertEqual(t, topElement, 15, "incorrect top element result")
}

func TestPopElementWhenExists(t *testing.T) {
	binaryHeap := NewBinaryHeap()
	binaryHeap.Push(15)
	binaryHeap.Push(11)
	binaryHeap.Push(4)
	binaryHeap.Push(3)
	binaryHeap.Push(5)
	binaryHeap.Push(25)
	var popElement int
	var err error
	if popElement, err = binaryHeap.Pop(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(popElement)
	}
	assertEqual(t, popElement, 25, "incorrect top element result")
	assertEqual(t, binaryHeap.Size(), 5, "incorrect binary heap size after pop")
}

func TestHeapSort(t *testing.T) {
	arr := []int{15,11,3,14,17,8,7,15,20,100,10,11}
	arrSortedByHeapSort := HeapSort(arr)
	arrSortedByStd := make([]int, len(arr))
	copy(arrSortedByStd, arr)
	sort.Ints(arrSortedByStd)
	assertEqual(t, arrSortedByHeapSort, arrSortedByStd, "heap sort implementation not working")
}
