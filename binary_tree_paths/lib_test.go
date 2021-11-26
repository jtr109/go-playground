package binarytreepaths

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(5),
	})
	expected := []string{"1->2->5", "1->3"}
	assert.Equal(t, expected, binaryTreePaths(root))
}
