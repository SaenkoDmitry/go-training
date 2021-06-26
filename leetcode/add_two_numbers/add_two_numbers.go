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

func AddTwoNumbers(first *ListNode, second *ListNode) *ListNode {

	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	root, accum := appendToResultSumTwoNode(nil, first, second, 0)
	last := root
	first = first.Next
	second = second.Next

	for first != nil && second != nil {
		last, accum = appendToResultSumTwoNode(last, first, second, accum)
		first, second, last = first.Next, second.Next, last.Next
	}
	for first != nil {
		last, accum = appendToResultOneNode(last, first, accum)
		first, last = first.Next, last.Next
	}
	for second != nil {
		last, accum = appendToResultOneNode(last, second, accum)
		second, last = second.Next, last.Next
	}
	if accum != 0 {
		last.Next = &ListNode{
			Val: accum,
		}
	}
	return root
}

func appendToResultSumTwoNode(last, firstCurr, secondCurr *ListNode, accum int) (*ListNode, int) {
	newNode := &ListNode{Val: (firstCurr.Val + secondCurr.Val + accum) % 10, Next: nil}
	accum = (firstCurr.Val + secondCurr.Val + accum) / 10
	if last == nil {
		last = newNode
	} else {
		last.Next = newNode
	}
	return last, accum
}

func appendToResultOneNode(last, curr *ListNode, accum int) (*ListNode, int) {
	newNode := &ListNode{Val: (curr.Val + accum) % 10, Next: nil}
	accum = (curr.Val + accum) / 10
	if last == nil {
		last = newNode
	} else {
		last.Next = newNode
	}
	return last, accum
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
