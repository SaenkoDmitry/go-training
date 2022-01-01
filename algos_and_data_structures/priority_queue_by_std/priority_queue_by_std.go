package priority_queue_by_std

import "container/heap"

/**
 * Definition for singly-linked list.
 * type v struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[0 : len(*pq)-1]
	return item
}

// TODO : how implement priority queue in go

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// ... Push(interface{]) items
}
