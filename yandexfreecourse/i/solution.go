package main

import "fmt"

// Comment it before submitting
type Node struct {
	data string
	Next *Node
}

func (n *Node) String(depth int) string {
	if n == nil || depth == 0 {
		return ""
	}
	return n.data + " " + n.Next.String(depth-1)
}

func Reverse(head *Node, left int, right int) *Node {
	index := 1
	curr := head
	for index < left-1 {
		curr = curr.Next
		index++
	}
	temp := reverseBeginningOfList(curr.Next, right-left)
	curr.Next = temp
	return head
}

func reverseBeginningOfList(head *Node, diff int) *Node {
	curr := head
	var newHead, prev *Node
	for curr != nil && diff >= 0 {
		newHead = curr
		curr = curr.Next
		newHead.Next = prev
		prev = newHead
		diff--
	}
	curr2 := newHead
	for curr2.Next != nil {
		curr2 = curr2.Next
	}
	curr2.Next = curr
	return newHead
}

func main() {
	input := &Node{
		data: "1",
		Next: &Node{
			data: "2",
			Next: &Node{
				data: "3",
				Next: &Node{
					data: "4",
					Next: &Node{
						data: "5",
					},
				},
			},
		},
	}
	res := Reverse(input, 2, 5)

	fmt.Println(res.String(10))

	input2 := &Node{
		data: "1",
		Next: &Node{
			data: "2",
		},
	}
	res2 := Reverse(input2, 0, 1)

	fmt.Println(res2.String(10))
}
