// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

package populatingnextrightpointersineachnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	currentLevel := []*Node{root}
	for len(currentLevel) != 0 {
		nextLevel := []*Node{}
		for i := 0; i < len(currentLevel)-1; i++ {
			currentLevel[i].Next = currentLevel[i+1]
		}
		for _, node := range currentLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return root
}
