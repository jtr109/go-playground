// https://leetcode.com/problems/binary-tree-right-side-view/

package binarytreerightsideview

import (
	"testing"

	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

type NilInt = treenode.NilInt

var NewInt = treenode.NewInt
var NewNil = treenode.NewNil

func TestExample1(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]NilInt{
		NewInt(1),
		NewInt(2),
		NewInt(3),
		NewNil(),
		NewInt(5),
		NewNil(),
		NewInt(4),
	})
	expected := []int{1, 3, 4}
	actual := rightSideView(op.Root())
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	op := treenode.NewOperator()
	op.FromSlice([]NilInt{})
	expected := []int{}
	actual := rightSideView(op.Root())
	assert.Equal(t, expected, actual)
}
