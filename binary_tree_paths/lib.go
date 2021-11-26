// https://programmercarl.com/0257.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84.html#%E9%80%92%E5%BD%92
//
// https://leetcode.com/problems/binary-tree-paths/

package binarytreepaths

import (
	"strconv"
	"strings"

	"github.com/jtr109/lcutils/treenode"
)

type TreeNode = treenode.TreeNode

func binaryTreePaths(root *TreeNode) (result []string) {
	p := binaryTreePathsInt(root)
	for _, paths := range p {
		result = append(result, strings.Join(paths, "->"))
	}
	return
}

func binaryTreePathsInt(root *TreeNode) (result [][]string) {
	if root == nil {
		return // ignore
	}
	if root.Left == nil && root.Right == nil {
		return [][]string{{strconv.Itoa(root.Val)}}
	}
	for _, leftPaths := range binaryTreePathsInt(root.Left) {
		r := []string{strconv.Itoa(root.Val)}
		r = append(r, leftPaths...)
		result = append(result, r)
	}
	for _, rightPaths := range binaryTreePathsInt(root.Right) {
		r := []string{strconv.Itoa(root.Val)}
		r = append(r, rightPaths...)
		result = append(result, r)
	}
	return
}
