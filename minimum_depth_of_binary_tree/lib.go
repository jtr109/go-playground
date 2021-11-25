// https://leetcode.com/problems/minimum-depth-of-binary-tree/

package minimumdepthofbinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	currentLevel := []*TreeNode{root}
outer:
	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range currentLevel {
			if node.Left == nil && node.Right == nil {
				break outer
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
		depth++
	}
	return depth
}
