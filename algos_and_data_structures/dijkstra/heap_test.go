package dijkstra

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

func Test_NewMinHeap(t *testing.T) {
	h := NewMinHeap()
	heap.Push(h, &Node{key: "A", cost: 1})
	heap.Push(h, &Node{key: "B", cost: 10})
	heap.Push(h, &Node{key: "C", cost: 7})
	heap.Push(h, &Node{key: "D", cost: 6})
	heap.Push(h, &Node{key: "E", cost: 5})
	heap.Push(h, &Node{key: "F", cost: 5})
	heap.Push(h, &Node{key: "G", cost: 4})
	heap.Push(h, &Node{key: "H", cost: 15})
	heap.Push(h, &Node{key: "I", cost: 13})
	got := make([]int, 0, h.Len())
	for h.Len() > 0 {
		item := heap.Pop(h).(*Node)
		got = append(got, item.cost)
		fmt.Println(item.key, item.cost)
	}
	expected := []int{1,4,5,5,6,7,10,13,15}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}
