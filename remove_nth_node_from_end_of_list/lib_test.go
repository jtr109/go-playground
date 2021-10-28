package removenthnodefromendoflist

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	n := 2
	expected := []int{1, 2, 3, 5}
	assert.Equal(t, expected, listnode.ConvertListNodeToArray(removeNthFromEnd(listnode.ConvertArrayToListNode(head), n)))
}

func TestExample2(t *testing.T) {
	head := []int{1}
	n := 1
	expected := []int{}
	assert.Equal(t, expected, listnode.ConvertListNodeToArray(removeNthFromEnd(listnode.ConvertArrayToListNode(head), n)))
}

func TestExample3(t *testing.T) {
	head := []int{1, 2}
	n := 1
	expected := []int{1}
	assert.Equal(t, expected, listnode.ConvertListNodeToArray(removeNthFromEnd(listnode.ConvertArrayToListNode(head), n)))
}
