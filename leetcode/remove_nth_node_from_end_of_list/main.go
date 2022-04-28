package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* идея в том, что нам надо уметь понимать, начали ли мы двигать окно или нет
* если не начали, то надо меня сразу head либо в nil, либо в head.Next
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var first, second *ListNode
	second = head
	for i := 0; i < n; i++ {
		second = second.Next
	}

	if second != nil {
		first = head
		for second != nil && second.Next != nil {
			first = first.Next
			second = second.Next
		}
	}

	if first != nil && first.Next != nil {
		first.Next = first.Next.Next
	}

	if first == nil {
		if head.Next == nil {
			head = nil
		} else {
			head = head.Next
		}
	}

	return head
}
