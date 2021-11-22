package invertbinarytree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

type NilInt = nilint.NilInt

var NewInt = nilint.NewInt
var NewNil = nilint.NewNil

func assertWorks(t *testing.T, root *[]NilInt, expected *[]NilInt) {
	op := treenode.NewOperator()
	op.FromSlice(*root)
	actualOperator := treenode.NewOperator()
	actualOperator.SetRoot(invertTree(op.Root()))
	assert.Equal(t, *expected, actualOperator.ToSlice())
}

func TestExample1(t *testing.T) {
	root := []NilInt{
		NewInt(4),
		NewInt(2),
		NewInt(7),
		NewInt(1),
		NewInt(3),
		NewInt(6),
		NewInt(9),
	}
	expected := []NilInt{
		NewInt(4),
		NewInt(7),
		NewInt(2),
		NewInt(9),
		NewInt(6),
		NewInt(3),
		NewInt(1),
	}
	assertWorks(t, &root, &expected)
}

func TestExample2(t *testing.T) {
	root := []NilInt{
		NewInt(2),
		NewInt(1),
		NewInt(3),
	}
	expected := []NilInt{
		NewInt(2),
		NewInt(3),
		NewInt(1),
	}
	assertWorks(t, &root, &expected)
}

func TestExample3(t *testing.T) {
	root := []NilInt{}
	expected := []NilInt{}
	assertWorks(t, &root, &expected)
}
