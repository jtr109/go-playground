package sametree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	p := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
	})
	q := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
	})
	assert.True(t, isSameTree(p, q))
}

func TestExample2(t *testing.T) {
	p := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
	})
	q := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewNil(),
		nilint.NewInt(2),
	})
	assert.False(t, isSameTree(p, q))
}

func TestExample3(t *testing.T) {
	p := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(1),
	})
	q := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(1),
		nilint.NewInt(2),
	})
	assert.False(t, isSameTree(p, q))
}
