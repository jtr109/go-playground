package binarytreezagziglevelordertraversal

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

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
	expected := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestExample2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
	})
	expected := [][]int{{1}}
	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestExample3(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{})
	expected := [][]int{}
	assert.Equal(t, expected, zigzagLevelOrder(root))
}
