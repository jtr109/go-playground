// https://leetcode.com/problems/remove-linked-list-elements/

package removelinkedlistelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
