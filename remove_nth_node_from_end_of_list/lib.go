// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package removenthnodefromendoflist

import "github.com/jtr109/lcutils/listnode"

type ListNode = listnode.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	virtualHead := &ListNode{
		Next: head,
	}
	previous := virtualHead
	current := virtualHead
	for i := 0; i < n; i++ {
		current = current.Next
	}
	for current.Next != nil {
		previous = previous.Next
		current = current.Next
	}
	previous.Next = previous.Next.Next
	return virtualHead.Next
}

// time complexity: O(N)
// space complexity: O(1)
