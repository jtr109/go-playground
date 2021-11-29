// https://leetcode.com/problems/search-in-a-binary-search-tree/

package searchinabinarysearchtree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	} else if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
