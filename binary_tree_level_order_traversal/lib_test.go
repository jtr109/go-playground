package binarytreelevelordertraversal

import (
	"testing"

	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]treenode.NilInt{
		treenode.NewInt(3),
		treenode.NewInt(9),
		treenode.NewInt(20),
		treenode.NewNil(),
		treenode.NewNil(),
		treenode.NewInt(15),
		treenode.NewInt(7),
	})
	root := op.Root()
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)
	assert.Equal(t, expected, actual)
}
