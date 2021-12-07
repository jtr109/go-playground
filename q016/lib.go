package q016

import (
	"github.com/jtr109/lcutils/listnode"
)

type ListNode = listnode.ListNode

func reverseGroup(root *ListNode, k int) *ListNode {
	cur := offset(root, length(root)%k-1)
	for cur.Next != nil {
		cur.Next = limitedReverse(cur.Next, k)
		cur = offset(cur, k)
	}
	return root
}

func length(root *ListNode) int {
	l := 0
	for cur := root; cur != nil; cur = cur.Next {
		l++
	}
	return l
}

func offset(root *ListNode, n int) *ListNode {
	cur := root
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	return cur
}

func limitedReverse(root *ListNode, limit int) *ListNode {
	cur := root
	var newRoot *ListNode
	for i := 0; i < limit; i++ {
		tmp := cur
		cur = cur.Next
		tmp.Next = newRoot
		newRoot = tmp
	}
	root.Next = cur
	return newRoot
}
