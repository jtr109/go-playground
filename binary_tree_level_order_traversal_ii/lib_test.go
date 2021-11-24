package binarytreelevelordertraversalii

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(9),
		nilint.NewInt(20),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(15),
		nilint.NewInt(7),
	})
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)
	assert.Equal(t, expected, actual)
}

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(9),
		nilint.NewInt(20),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(15),
		nilint.NewInt(7),
	})
	expected := [][]int{{15, 7}, {9, 20}, {3}}
	actual := levelOrderBottom(root)
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
	})
	actual := levelOrderBottom(root)
	expected := [][]int{{1}}
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{})
	actual := levelOrderBottom(root)
	expected := [][]int{}
	assert.Equal(t, expected, actual)
}
