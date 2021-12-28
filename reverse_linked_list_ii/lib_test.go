package reverselinkedlistii

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	root := listnode.FromSlice([]int{1, 2, 3, 4})
	actual := reverse(root)
	assert.Equal(t, 4, actual.Val)
	assert.Equal(t, 3, actual.Next.Val)
	assert.Equal(t, 2, actual.Next.Next.Val)
	assert.Equal(t, 1, actual.Next.Next.Next.Val)
}
