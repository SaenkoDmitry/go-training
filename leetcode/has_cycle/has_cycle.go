package has_cycle

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
