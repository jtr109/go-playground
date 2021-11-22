// https://leetcode.com/problems/invert-binary-tree/

package invertbinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
