// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

package minimumabsolutedifferenceinbst

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func getMinimumDifference(root *TreeNode) int {
	s := inorderTraversal(root)
	return minDiff(s)
}

func inorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return []int{}
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDiff(s []int) int {
	diff := s[1] - s[0]
	for i := 1; i < len(s)-1; i++ {
		diff = min(diff, s[i+1]-s[i])
	}
	return diff
}
