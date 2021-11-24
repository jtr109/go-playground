package removenthnodefromendoflist

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func isWorks(t *testing.T, head []int, n int, expected []int) {
	assert.Equal(
		t,
		listnode.NewOperator().FromSlice(expected).Head(),
		removeNthFromEnd(listnode.NewOperator().FromSlice(head).Head(), n),
	)
}

func TestExample1(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	n := 2
	expected := []int{1, 2, 3, 5}
	isWorks(t, head, n, expected)
}

func TestExample2(t *testing.T) {
	head := []int{1}
	n := 1
	expected := []int{}
	isWorks(t, head, n, expected)
}

func TestExample3(t *testing.T) {
	head := []int{1, 2}
	n := 1
	expected := []int{1}
	isWorks(t, head, n, expected)
}
