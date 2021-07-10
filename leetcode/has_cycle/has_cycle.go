package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fastCurr := head
	slowCurr := head

	for {
		if fastCurr.Next != nil && fastCurr.Next.Next != nil {
			fastCurr = fastCurr.Next.Next
		} else {
			break
		}

		if slowCurr.Next != nil {
			slowCurr = slowCurr.Next
		}

		if fastCurr == slowCurr {
			return true
		}
	}
	return false
}

func main() {
	secondListNode := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: -4,
				Next: nil,
			},
		},
	}
	secondListNode.Next.Next.Next = secondListNode
	cycledlist := &ListNode{
		Val: 3,
		Next: secondListNode,
	}
	fmt.Println(hasCycle(cycledlist))
}
