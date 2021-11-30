// https://leetcode.com/problems/delete-node-in-a-bst/

package deletenodeinabst

import "github.com/jtr109/lcutils/treenode"

type TreeNode = treenode.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil && root.Right == nil { // the root is a leaf
		return nil
	} else if root.Left != nil && root.Right != nil { // the root has both left and right children
		current := root.Right
		for current.Left != nil {
			current = current.Left
		}
		current.Left = root.Left
		return root.Right
	} else if root.Left == nil { // the root only has a right child
		return root.Right
	} else {
		return root.Left
	}
}
