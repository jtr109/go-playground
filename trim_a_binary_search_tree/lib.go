// https://leetcode.com/problems/trim-a-binary-search-tree/

package trimabinarysearchtree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
