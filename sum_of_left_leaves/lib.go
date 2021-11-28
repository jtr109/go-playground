// https://leetcode.com/problems/sum-of-left-leaves/

package sumofleftleaves

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func sumOfLeftLeaves(root *TreeNode) (result int) {
	if root == nil {
		return 0
	}
	result += sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		result += root.Left.Val
	}
	return
}
