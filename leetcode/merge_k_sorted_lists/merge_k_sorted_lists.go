package main

import (
	"container/heap"
	"fmt"
	pqlistnode "github.com/SaenkoDmitry/go-training/algos_and_data_structures/priority_queue_by_std"
)

func mergeKLists(lists []*pqlistnode.ListNode) *pqlistnode.ListNode {
	pq := make(pqlistnode.PriorityQueue, 0)
	heap.Init(&pq)
	for i := range lists {
		heap.Push(&pq, lists[i])
	}

	var root, curr *pqlistnode.ListNode
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*pqlistnode.ListNode)
		newElem := &pqlistnode.ListNode{Val: node.Val}
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

func main() {
	list1 := &pqlistnode.ListNode{Val: 1, Next: &pqlistnode.ListNode{Val: 4, Next: &pqlistnode.ListNode{Val: 5}}}
	list2 := &pqlistnode.ListNode{Val: 1, Next: &pqlistnode.ListNode{Val: 3, Next: &pqlistnode.ListNode{Val: 4}}}
	list3 := &pqlistnode.ListNode{Val: 2, Next: &pqlistnode.ListNode{Val: 6}}
	fmt.Println(mergeKLists([]*pqlistnode.ListNode{list1, list2, list3}))
}
