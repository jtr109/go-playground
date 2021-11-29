package lowestcommonancestorofabinarytree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample2(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(3),
		nilint.NewInt(5),
		nilint.NewInt(1),
		nilint.NewInt(6),
		nilint.NewInt(2),
		nilint.NewInt(0),
		nilint.NewInt(8),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(7),
		nilint.NewInt(4),
	})
	p := root.Left
	q := root.Left.Right.Right
	expected := root.Left
	assert.Equal(t, expected, lowestCommonAncestor(root, p, q))
}

func TestSubmission1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{nilint.NewInt(1), nilint.NewInt(2)})
	p := root.Left
	q := root
	expected := root
	assert.Equal(t, expected, lowestCommonAncestor(root, p, q))
}
