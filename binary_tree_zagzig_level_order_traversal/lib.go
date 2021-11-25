// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

package binarytreezagziglevelordertraversal

import (
	"github.com/jtr109/lcutils/treenode"
)

type TreeNode = treenode.TreeNode

func zigzagLevelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return [][]int{}
	}

	currentLevel := []*TreeNode{root}
	level := 0
	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		currentLevelValue := []int{}
		for _, node := range currentLevel {
			currentLevelValue = append(currentLevelValue, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if level%2 == 1 {
			for i, j := 0, len(currentLevelValue)-1; i < j; {
				currentLevelValue[i], currentLevelValue[j] = currentLevelValue[j], currentLevelValue[i]
				i++
				j--
			}
		}
		result = append(result, currentLevelValue)
		currentLevel = nextLevel
		level++
	}
	return
}
