package reverselinkedlistii

import (
	"github.com/jtr109/lcutils/listnode"
)

type ListNode = listnode.ListNode

func reverse(root *ListNode) *ListNode {
	current := root
	var prev *ListNode
	for current != nil {
		tmp := current
		current = current.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}
