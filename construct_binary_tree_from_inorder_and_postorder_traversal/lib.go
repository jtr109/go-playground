// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package constructbinarytreefrominorderandpostordertraversal

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	i := 0
	for inorder[i] != val {
		i++
	}
	inorderLeft, inorderRight := inorder[:i], inorder[i+1:]
	postorderLeft, postorderRight := postorder[:i], postorder[i:len(postorder)-1]
	return &TreeNode{
		Val:   val,
		Left:  buildTree(inorderLeft, postorderLeft),
		Right: buildTree(inorderRight, postorderRight),
	}
}
