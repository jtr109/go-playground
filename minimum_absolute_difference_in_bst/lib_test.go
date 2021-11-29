package minimumabsolutedifferenceinbst

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(4),
		nilint.NewInt(2),
		nilint.NewInt(6),
		nilint.NewInt(1),
		nilint.NewInt(3),
	})
	expected := 1
	assert.Equal(t, expected, getMinimumDifference(root))
}

func TestSubmission1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewNil(),
		nilint.NewInt(3),
		nilint.NewInt(2),
	})
	expected := 1
	assert.Equal(t, expected, getMinimumDifference(root))
}

func TestSubmission2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(236),
		nilint.NewInt(104),
		nilint.NewInt(701),
		nilint.NewNil(),
		nilint.NewInt(227),
		nilint.NewNil(),
		nilint.NewInt(911),
	})
	expected := 9
	assert.Equal(t, expected, getMinimumDifference(root))
}
