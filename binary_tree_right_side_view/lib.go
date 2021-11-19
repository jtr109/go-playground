package binarytreerightsideview

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func rightSideView(root *TreeNode) (result []int) {
	if root == nil {
		return []int{}
	}
	currentLevelQueue := []*TreeNode{root}
	for len(currentLevelQueue) != 0 {
		result = append(result, currentLevelQueue[len(currentLevelQueue)-1].Val)
		nextLevelQueue := []*TreeNode{}
		for _, node := range currentLevelQueue {
			if node.Left != nil {
				nextLevelQueue = append(nextLevelQueue, node.Left)
			}
			if node.Right != nil {
				nextLevelQueue = append(nextLevelQueue, node.Right)
			}
		}
		currentLevelQueue = nextLevelQueue
	}
	return
}
