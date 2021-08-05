// https://leetcode.com/explore/featured/card/august-leetcoding-challenge-2021/613/week-1-august-1st-august-7th/3838/

package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	newTargetSum := targetSum - root.Val
	leftPathSum := [][]int{}
	res := [][]int{}
	if root.Left != nil {
		leftPathSum = pathSum(root.Left, newTargetSum)
	}
	rightPathSum := [][]int{}
	if root.Right != nil {
		rightPathSum = pathSum(root.Right, newTargetSum)
	}
	for _, ps := range leftPathSum {
		res = append(res, append([]int{root.Val}, ps...))
	}
	for _, ps := range rightPathSum {
		res = append(res, append([]int{root.Val}, ps...))
	}
	return res
}
