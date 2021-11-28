package pathsum

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(5),
		nilint.NewInt(4),
		nilint.NewInt(8),
		nilint.NewInt(11),
		nilint.NewNil(),
		nilint.NewInt(13),
		nilint.NewInt(4),
		nilint.NewInt(7),
		nilint.NewInt(2),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(1),
	})
	targetSum := 22
	assert.True(t, hasPathSum(root, targetSum))
}

func TestSubmission1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{})
	targetSum := 0
	assert.False(t, hasPathSum(root, targetSum))
}
