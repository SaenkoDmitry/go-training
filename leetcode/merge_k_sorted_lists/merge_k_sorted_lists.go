package main

import (
	"container/heap"
	"fmt"
	"strconv"
)


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

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := range lists {
		heap.Push(&pq, lists[i])
	}

	var root, curr *ListNode
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		newElem := &ListNode{Val: node.Val}
		if root == nil {
			root = newElem
			curr = root
		} else {
			curr.Next = newElem
			curr = newElem
		}
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return root
}

func (l ListNode) String() string {
	currItemStr := strconv.FormatInt(int64(l.Val), 10)
	if l.Next == nil {
		return currItemStr
	}
	return currItemStr + " " + l.Next.String()
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	fmt.Println(mergeKLists([]*ListNode{list1, list2, list3}))
}
