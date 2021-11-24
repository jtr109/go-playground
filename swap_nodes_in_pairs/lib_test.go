package swapnodesinpairs

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	head := []int{1, 2, 3, 4}
	expected := []int{2, 1, 4, 3}
	assert.Equal(
		t,
		listnode.FromSlice(expected),
		swapPairs(listnode.FromSlice(head)),
	)
}
