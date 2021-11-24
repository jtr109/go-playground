// https://leetcode.com/problems/maximum-depth-of-binary-tree/

package maximumdepthofbinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	} else {
		return rightMaxDepth + 1
	}
}
