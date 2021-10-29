// https://leetcode.com/problems/linked-list-cycle-ii/

package linkedlistcycleii

import "github.com/jtr109/lcutils/listnode"

type ListNode = listnode.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			index1 := head
			index2 := slow // meeting point
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}
