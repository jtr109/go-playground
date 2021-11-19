package binarytreelevelordertraversal

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

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
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)
	assert.Equal(t, expected, actual)
}
