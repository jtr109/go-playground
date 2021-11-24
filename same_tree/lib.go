// https://leetcode.com/problems/same-tree/

package sametree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
	for len(queue) >= 2 {
		currentP, currentQ := queue[0], queue[1]
		queue = queue[2:]
		if currentP == nil && currentQ == nil {
			continue
		}
		if currentP == nil || currentQ == nil { // only 1 node is nil
			return false
		}
		if currentP.Val != currentQ.Val {
			return false
		}
		queue = append(queue, currentP.Left, currentQ.Left, currentP.Right, currentQ.Right)
	}
	return len(queue) == 0
}
