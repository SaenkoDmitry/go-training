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

func (head ListNode) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(int64(head.Val), 10))
	curr := head.Next
	for curr != nil {
		buffer.WriteRune(' ')
		buffer.WriteString(strconv.FormatInt(int64(curr.Val), 10))
		curr = curr.Next
	}
	return buffer.String()
}

func (head *ListNode) Reverse() *ListNode {
	var root, prevRoot, temp *ListNode
	curr := head

	for curr != nil {
		temp = curr.Next

		root = curr
		root.Next = prevRoot
		prevRoot = root

		curr = temp
	}
	return root
}

func ArrayToListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNode{
		Val: arr[0],
	}
	last := root
	var newNode *ListNode
	for i := 1; i < len(arr); i++ {
		newNode = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		last.Next = newNode
		last = last.Next
	}
	return root
}

func main() {
	first := ArrayToListNode([]int{2, 4, 3})
	fmt.Println(first)
	fmt.Println(first.Reverse())
}
