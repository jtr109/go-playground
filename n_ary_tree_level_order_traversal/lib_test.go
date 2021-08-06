package narytreelevelordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	root := Node{
		1,
		[]*Node{
			{
				3,
				[]*Node{
					{5, []*Node{}},
					{6, []*Node{}},
				},
			},
			{2, []*Node{}},
			{4, []*Node{}},
		},
	}
	assert.Equal(t, [][]int{{1}, {3, 2, 4}, {5, 6}}, levelOrder(&root))
}
