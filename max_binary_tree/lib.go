// https://leetcode.com/problems/maximum-binary-tree/

package maxbinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := 0
	for i, n := range nums {
		if n > nums[maxIndex] {
			maxIndex = i
		}
	}
	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}

// time complexity: N * logN
