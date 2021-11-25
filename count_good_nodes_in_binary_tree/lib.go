// https://leetcode.com/problems/count-good-nodes-in-binary-tree/

package countgoodnodesinbinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func goodNodes(root *TreeNode) int {
	return countGoodNodes(root, root.Val)
}

func countGoodNodes(node *TreeNode, maxInPath int) (count int) {
	if node == nil {
		return 0
	}
	if node.Val >= maxInPath {
		count++
		maxInPath = node.Val
	}
	count += countGoodNodes(node.Left, maxInPath) + countGoodNodes(node.Right, maxInPath)
	return
}
