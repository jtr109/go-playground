package sumofleftleaves

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
	expected := 24
	assert.Equal(t, expected, sumOfLeftLeaves(root))
}

func TestExample2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{nilint.NewInt(1)})
	expected := 0
	assert.Equal(t, expected, sumOfLeftLeaves(root))
}

func TestSubmission1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
		nilint.NewInt(4),
		nilint.NewInt(5),
	})
	expected := 4
	assert.Equal(t, expected, sumOfLeftLeaves(root))
}
