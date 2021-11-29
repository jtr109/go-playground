package lowestcommonancestorofabinarytree

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	commonAncestor, _ := lowestCommonAncestorWithCount(root, p, q)
	return commonAncestor
}

func lowestCommonAncestorWithCount(root, p, q *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	posteritiesCount := 0
	if root == p {
		posteritiesCount++
	}
	if root == q {
		posteritiesCount++
	}
	if posteritiesCount == 2 { // root is both p and q
		return root, posteritiesCount
	}
	leftAncestor, leftPosteritiesCount := lowestCommonAncestorWithCount(root.Left, p, q)
	if leftPosteritiesCount == 2 { // root.Left is the ancestor of p and q
		return leftAncestor, leftPosteritiesCount
	}
	posteritiesCount += leftPosteritiesCount
	if posteritiesCount == 2 { // both p and q are posterities of either root.Left or root
		return root, posteritiesCount
	}
	rightAncestor, rightPosteritiesCount := lowestCommonAncestorWithCount(root.Right, p, q)
	if rightPosteritiesCount == 2 { // root.Right is the ancestor of p and q
		return rightAncestor, rightPosteritiesCount
	}
	posteritiesCount += rightPosteritiesCount
	if posteritiesCount == 2 { // both p and q are posterities of either root
		return root, posteritiesCount
	}
	return nil, posteritiesCount
}
