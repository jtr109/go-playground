package symmetrictree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := []nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(2),
		nilint.NewInt(3),
		nilint.NewInt(4),
		nilint.NewInt(4),
		nilint.NewInt(3),
	}

	assert.True(t, isSymmetric(treenode.FromSlice(root)))
}

func TestExample2(t *testing.T) {
	root := []nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(2),
		nilint.NewNil(),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(3),
	}
	treenode.FromSlice(root)
	assert.False(t, isSymmetric(treenode.FromSlice(root)))
}
