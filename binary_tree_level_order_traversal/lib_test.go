package binarytreelevelordertraversal

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
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)
	assert.Equal(t, expected, actual)
}
