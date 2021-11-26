// https://programmercarl.com/0110.%E5%B9%B3%E8%A1%A1%E4%BA%8C%E5%8F%89%E6%A0%91.html#%E9%80%92%E5%BD%92
//
// https://leetcode.com/problems/balanced-binary-tree/

package balancedbinarytree

import (
	"github.com/jtr109/lcutils/treenode"
)

type TreeNode = treenode.TreeNode

func isBalanced(root *TreeNode) bool {
	_, b := height(root)
	return b
}

func height(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	hl, b := height(root.Left)
	if !b {
		return 0, false
	}
	hr, b := height(root.Right)
	if !b {
		return 0, false
	}
	if hl == hr {
		return hl + 1, true
	}
	if hl > hr && hl-hr <= 1 {
		return hl + 1, true
	}
	if hr > hl && hr-hl <= 1 {
		return hr + 1, true
	}
	return 0, false
}
