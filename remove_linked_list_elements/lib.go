package removelinkedlistelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var last, current *ListNode = nil, head
	for current != nil {
		if current.Val == val {
			last.Next = current.Next
			current = current.Next
		} else {
			last = current
			current = current.Next
		}
	}
	return head
}
