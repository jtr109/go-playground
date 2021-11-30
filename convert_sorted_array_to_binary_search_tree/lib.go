// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

package convertsortedarraytobinarysearchtree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	return traversal(nums, 0, len(nums)-1)
}

func traversal(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  traversal(nums, left, mid-1),
		Right: traversal(nums, mid+1, right),
	}
}
