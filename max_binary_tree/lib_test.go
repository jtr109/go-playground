package maxbinarytree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	expected := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(6),
		nilint.NewInt(3),
		nilint.NewInt(5),
		nilint.NewNil(),
		nilint.NewInt(2),
		nilint.NewInt(0),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(1),
	})
	assert.Equal(t, expected, constructMaximumBinaryTree(nums))
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 1}
	expected := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(2),
		nilint.NewNil(),
		nilint.NewInt(1),
	})
	assert.Equal(t, expected, constructMaximumBinaryTree(nums))
}
