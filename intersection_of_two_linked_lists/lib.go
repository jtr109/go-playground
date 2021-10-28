// https://leetcode.com/problems/intersection-of-two-linked-lists/

package intersectionoftwolinkedlists

import "github.com/jtr109/lcutils/listnode"

type ListNode = listnode.ListNode

// time complexity: O(N)
// space complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := lengthListNode(headA)
	lengthB := lengthListNode(headB)
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			headB = headB.Next
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func lengthListNode(head *ListNode) int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}
