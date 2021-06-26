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
	buffer.WriteString(strconv.FormatInt(int64(l.Val), 10))
	curr := l.Next
	for curr != nil {
		buffer.WriteRune(' ')
		buffer.WriteString(strconv.FormatInt(int64(curr.Val), 10))
		curr = curr.Next
	}
	return buffer.String()
}

func (l *ListNode) Reverse() *ListNode {
	curr := l
	var futureRoot, prevRoot, temp *ListNode
	for curr != nil {
		temp = curr.Next

		futureRoot = curr
		futureRoot.Next = prevRoot
		prevRoot = futureRoot

		curr = temp
	}
	return futureRoot
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
