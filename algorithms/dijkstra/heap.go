package dijkstra

import "container/heap"

type Node struct {
	key  string
	cost int
}

type MinHeap []*Node

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*Node))
}

func (m *MinHeap) Pop() interface{} {
	if m.Len() > 0 {
		tail := (*m)[len(*m)-1]
		*m = (*m)[:len(*m)-1]
		return tail
	}
	return struct{}{}
}

func NewMinHeap() *MinHeap {
	h := make(MinHeap, 0)
	heap.Init(&h)
	return &h
}
