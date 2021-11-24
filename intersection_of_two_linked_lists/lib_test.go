package intersectionoftwolinkedlists

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestLengthListNode(t *testing.T) {
	assert.Equal(t, 0, lengthListNode(listnode.FromSlice([]int{})))
	assert.Equal(t, 1, lengthListNode(listnode.FromSlice([]int{1})))
}

func TestExample1(t *testing.T) {
	headA := listnode.FromSlice([]int{4, 1, 8, 4, 5}) // 4, 1, 8, 4, 5
	headB := listnode.FromSlice([]int{5, 6, 1})       // 5, 6, 1, 8, 4, 5
	headB.Next.Next.Next = headA.Next.Next
	assert.Equal(t, 8, getIntersectionNode(headA, headB).Val)
}
