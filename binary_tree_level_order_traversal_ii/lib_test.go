package binarytreelevelordertraversalii

import (
	"testing"

	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
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
	expected := [][]int{{15, 7}, {9, 20}, {3}}
	actual := levelOrderBottom(root)
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]treenode.NilInt{
		treenode.NewInt(1),
	})
	actual := levelOrderBottom(op.Root())
	expected := [][]int{{1}}
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]treenode.NilInt{})
	actual := levelOrderBottom(op.Root())
	expected := [][]int{}
	assert.Equal(t, expected, actual)
}

func TestSubmission1(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]treenode.NilInt{
		treenode.NewInt(1),
		treenode.NewInt(2),
		treenode.NewNil(),
		treenode.NewInt(3),
		treenode.NewNil(),
		treenode.NewInt(4),
		treenode.NewNil(),
		treenode.NewInt(5),
	})
}
