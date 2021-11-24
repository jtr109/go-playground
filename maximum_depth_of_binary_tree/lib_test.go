package maximumdepthofbinarytree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.NewOperator().FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(9),
		nilint.NewInt(20),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(15),
		nilint.NewInt(7),
	}).Root()
	actual := maxDepth(root)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	root := treenode.NewOperator().FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewNil(),
		nilint.NewInt(2),
	}).Root()
	actual := maxDepth(root)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	root := treenode.NewOperator().FromSlice([]nilint.NilInt{}).Root()
	actual := maxDepth(root)
	expected := 0
	assert.Equal(t, expected, actual)
}
