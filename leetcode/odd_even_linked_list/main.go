package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(l.Val))
	for curr := l.Next; curr != nil; curr = curr.Next {
		buffer.WriteString(", ")
		buffer.WriteString(strconv.Itoa(curr.Val))
	}
	return buffer.String()
}

func oddEvenList(head *ListNode) *ListNode {
	var firstOdd *ListNode = nil
	var lastOdd *ListNode = nil
	var firstEven *ListNode = nil
	var lastEven *ListNode = nil

	var temp *ListNode
	for curr := head; curr != nil; {
		//fmt.Println("-", curr.Val)
		temp = curr.Next
		if curr.Val%2 == 1 {
			if firstOdd == nil {
				firstOdd = curr
			}
			if lastOdd != nil {
				lastOdd.Next = curr
			}
			lastOdd = curr
			lastOdd.Next = firstEven
		} else {
			if firstEven == nil {
				firstEven = curr
			}
			if lastEven != nil {
				lastEven.Next = curr
			}
			lastEven = curr
			lastEven.Next = nil
		}
		curr = temp
	}

	if firstOdd != nil {
		return firstOdd
	}
	return firstEven
}

func main() {
	list1 := &ListNode{2, &ListNode{2, &ListNode{1, nil}}}
	list2 := &ListNode{1, &ListNode{2, &ListNode{3,
		&ListNode{4, &ListNode{5, nil}},
	}}}
	list3 := &ListNode{2, &ListNode{1, &ListNode{3,
		&ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}},
	}}}
	fmt.Println(oddEvenList(list1))
	fmt.Println(oddEvenList(list2))
	fmt.Println(oddEvenList(list3))
}
