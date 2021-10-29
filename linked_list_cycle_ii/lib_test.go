package linkedlistcycleii

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	arr := []int{3, 2, 0, -4}
	head := listnode.ConvertArrayToListNode(arr)
	head.Next.Next.Next.Next = head.Next
	expected := 2
	assert.Equal(t, expected, detectCycle(head).Val)
}
