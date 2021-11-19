package binarytreelevelordertraversalii

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(9),
		nilint.NewInt(20),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(15),
		nilint.NewInt(7),
	})
	root := op.Root()
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)
	assert.Equal(t, expected, actual)
}

func TestExample1(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(9),
		nilint.NewInt(20),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(15),
		nilint.NewInt(7),
	})
	root := op.Root()
	expected := [][]int{{15, 7}, {9, 20}, {3}}
	actual := levelOrderBottom(root)
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
	})
	actual := levelOrderBottom(op.Root())
	expected := [][]int{{1}}
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]nilint.NilInt{})
	actual := levelOrderBottom(op.Root())
	expected := [][]int{}
	assert.Equal(t, expected, actual)
}

func TestSubmission1(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewNil(),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(4),
		nilint.NewNil(),
		nilint.NewInt(5),
	})
}
