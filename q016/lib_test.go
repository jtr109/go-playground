package q016

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	root := listnode.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	k := 3
	expected := []int{1, 2, 5, 4, 3, 8, 7, 6}
	for cur, i := reverseGroup(root, k), 0; cur != nil; {
		assert.Equal(t, expected[i], cur.Val)
		cur = cur.Next
		i++
	}
}

func TestLength(t *testing.T) {
	root := listnode.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	assert.Equal(t, 8, length(root))
}

func TestOffset(t *testing.T) {
	root := listnode.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	assert.Equal(t, 3, offset(root, 2).Val)
}

func TestLimitedReverse(t *testing.T) {
	root := listnode.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	actual := limitedReverse(root, 3)
	assert.Equal(t, 3, actual.Val)
	assert.Equal(t, 2, actual.Next.Val)
	assert.Equal(t, 1, actual.Next.Next.Val)
	assert.Equal(t, 4, actual.Next.Next.Next.Val)
}
