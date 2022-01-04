package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next  *ListNode
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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, curr, newElem *ListNode
	var rest int
	for l1 != nil || l2 != nil || rest != 0 {
		var first, second int
		if l1 != nil {
			first = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			second = l2.Val
			l2 = l2.Next
		}
		newElem = &ListNode{
			Val: (first + second + rest) % 10,
		}
		rest = (first + second + rest) / 10
		if result == nil {
			result = newElem
			curr = newElem
		} else {
			curr.Next = newElem
			curr = curr.Next
		}
	}
	return result
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
			Val: arr[i],
			Next: nil,
		}
		last.Next = newNode
		last = last.Next
	}
	return root
}

func main() {
	first := ArrayToListNode([]int{2,4,3})
	second := ArrayToListNode([]int{5,6,4})
	fmt.Println(AddTwoNumbers(first, second))

	first = ArrayToListNode([]int{9,9,9,9,9,9,9})
	second = ArrayToListNode([]int{9,9,9,9})
	fmt.Println(AddTwoNumbers(first, second))
}
