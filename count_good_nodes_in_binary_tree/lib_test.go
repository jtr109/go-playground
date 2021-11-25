package countgoodnodesinbinarytree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(1),
		nilint.NewInt(4),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(1),
		nilint.NewInt(5),
	})
	expected := 4
	assert.Equal(t, expected, goodNodes(root))
}

func TestExample2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(4),
		nilint.NewInt(2),
	})
	expected := 3
	assert.Equal(t, expected, goodNodes(root))
}

func TestExample3(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
	})
	expected := 1
	assert.Equal(t, expected, goodNodes(root))
}

func TestSubmission1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(-1),
		nilint.NewInt(5),
		nilint.NewInt(-2),
		nilint.NewInt(4),
		nilint.NewInt(4),
		nilint.NewInt(2),
		nilint.NewInt(-2),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(-4),
		nilint.NewNil(),
		nilint.NewInt(-2),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewInt(-2),
		nilint.NewInt(0),
		nilint.NewNil(),
		nilint.NewInt(-1),
		nilint.NewNil(),
		nilint.NewInt(-3),
		nilint.NewNil(),
		nilint.NewInt(-4),
		nilint.NewInt(-3),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(3),
		nilint.NewInt(-3),
	})
	expected := 5
	assert.Equal(t, expected, goodNodes(root))
}
