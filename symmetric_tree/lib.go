// https://leetcode.com/problems/symmetric-tree/

package symmetrictree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	}
	// no one of nodes is nil now
	if node1.Val != node2.Val {
		return false
	}
	if !isMirror(node1.Left, node2.Right) {
		return false
	}
	if !isMirror(node1.Right, node2.Left) {
		return false
	}
	return true
}
