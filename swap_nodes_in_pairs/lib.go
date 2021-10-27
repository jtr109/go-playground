// https://leetcode.com/problems/swap-nodes-in-pairs/

package swapnodesinpairs

import (
	"github.com/jtr109/lcutils/listnode"
)

func swapPairs(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	virtualHead := &listnode.ListNode{
		Next: head,
	}
	var previous *listnode.ListNode = virtualHead
	current := head
	for current != nil && current.Next != nil {
		// move current.Next before current
		previous.Next = current.Next
		current.Next = current.Next.Next
		previous.Next.Next = current
		// reset previous and current pointers
		previous = current
		current = current.Next
	}
	return virtualHead.Next
}

// time complexity: O(N)
// space complexity: O(1)
