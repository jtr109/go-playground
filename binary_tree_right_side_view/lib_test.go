// https://leetcode.com/problems/binary-tree-right-side-view/

package binarytreerightsideview

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

type NilInt = nilint.NilInt

var NewInt = nilint.NewInt
var NewNil = nilint.NewNil

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]NilInt{
		NewInt(1),
		NewInt(2),
		NewInt(3),
		NewNil(),
		NewInt(5),
		NewNil(),
		NewInt(4),
	})
	expected := []int{1, 3, 4}
	actual := rightSideView(root)
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	root := treenode.FromSlice([]NilInt{})
	expected := []int{}
	actual := rightSideView(root)
	assert.Equal(t, expected, actual)
}
