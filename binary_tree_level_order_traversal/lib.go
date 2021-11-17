// https://leetcode.com/problems/binary-tree-level-order-traversal/

package binarytreelevelordertraversal

import (
	"github.com/jtr109/lcutils/treenode"
)

type TreeNode = treenode.TreeNode

func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		layer := []int{}
		for _, node := range queue {
			layer = append(layer, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		result = append(result, layer)
		queue = nextQueue
	}
	return
}
