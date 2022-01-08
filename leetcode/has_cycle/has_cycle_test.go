package has_cycle

import "testing"

func Test_hasCycle(t *testing.T) {
	cycledlist := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  -4,
				Next: nil,
			},
		},
	}
	cycledlist.Next.Next.Next = cycledlist
	tests := []struct {
		name string
		args *ListNode
		want bool
	}{
		{
			"no cycle",
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  -4,
						Next: nil,
					},
				},
			},
			false,
		},
		{
			"has cycle",
			cycledlist,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
