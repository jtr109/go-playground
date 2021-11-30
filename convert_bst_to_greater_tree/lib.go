// https://leetcode.com/problems/convert-bst-to-greater-tree/

package convertbsttogreatertree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

// func convertBST(root *TreeNode) *TreeNode {
// 	result, _ := traversal(root, 0)
// 	return result
// }

// func traversal(root *TreeNode, sum int) (*TreeNode, int) {
// 	if root == nil {
// 		return root, sum
// 	}
// 	right, sum := traversal(root.Right, sum)
// 	sum += root.Val
// 	val := sum
// 	left, sum := traversal(root.Left, sum)
// 	return &TreeNode{
// 		Val:   val,
// 		Left:  left,
// 		Right: right,
// 	}, sum
// }

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
		}
	}
	return root
}
