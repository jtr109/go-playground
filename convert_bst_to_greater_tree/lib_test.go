package convertbsttogreatertree

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/jtr109/lcutils/treenode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := treenode.FromSlice([]nilint.NilInt{
		nilint.NewInt(4),
		nilint.NewInt(1),
		nilint.NewInt(6),
		nilint.NewInt(0),
		nilint.NewInt(2),
		nilint.NewInt(5),
		nilint.NewInt(7),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(3),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewNil(),
		nilint.NewInt(8),
	})
	assert.Equal(t, nil, convertBST(root))
}
