// https://leetcode.com/problems/path-sum/

package pathsum

import (
	"github.com/jtr109/lcutils/treenode"
)

type TreeNode = treenode.TreeNode

func hasPathSum(root *TreeNode, targetSum int) (result bool) {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil { // leaf node
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
