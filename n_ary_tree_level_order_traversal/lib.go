// https://leetcode.com/explore/featured/card/august-leetcoding-challenge-2021/613/week-1-august-1st-august-7th/3871/

package narytreelevelordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	childrenLevelOrder := [][]int{{root.Val}}
	for _, child := range root.Children {
		childLevelOrder := levelOrder(child)
		for i, order := range childLevelOrder {
			if i+1 >= len(childrenLevelOrder) {
				childrenLevelOrder = append(childrenLevelOrder, []int{})
			}
			childrenLevelOrder[i+1] = append(childrenLevelOrder[i+1], order...)
		}
	}
	return childrenLevelOrder
}
