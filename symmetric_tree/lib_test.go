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
	op := treenode.NewOperator()
	op.FromSlice(root)
	assert.True(t, isSymmetric(op.Root()))
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
	op := treenode.NewOperator()
	op.FromSlice(root)
	assert.False(t, isSymmetric(op.Root()))
}
