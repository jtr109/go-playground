// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

package binarytreelevelordertraversalii

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	reverseSlice(result)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		currentLevel := []int{}
		for _, node := range queue {
			currentLevel = append(currentLevel, node.Val)
			left, right := node.Left, node.Right
			if left != nil {
				nextQueue = append(nextQueue, left)
			}
			if right != nil {
				nextQueue = append(nextQueue, right)
			}
		}
		result = append(result, currentLevel)
		queue = nextQueue
	}
	return result
}

func reverseSlice(s [][]int) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
