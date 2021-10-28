package intersectionoftwolinkedlists

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestLengthListNode(t *testing.T) {
	assert.Equal(t, 0, lengthListNode(listnode.ConvertArrayToListNode([]int{})))
	assert.Equal(t, 1, lengthListNode(listnode.ConvertArrayToListNode([]int{1})))
}

func TestExample1(t *testing.T) {
	headA := listnode.ConvertArrayToListNode([]int{4, 1, 8, 4, 5})
	headB := listnode.ConvertArrayToListNode([]int{5, 6, 1, 8, 4, 5})
	assert.Equal(t, 8, getIntersectionNode(headA, headB).Val)
}
