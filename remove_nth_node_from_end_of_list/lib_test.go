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
	actual, err := listnode.NewOperator().SetHead(removeNthFromEnd(listnode.NewOperator().FromSlice(head).Head(), n)).ToSlice()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	head := []int{1}
	n := 1
	expected := []int{}
	actual, err := listnode.NewOperator().SetHead(removeNthFromEnd(listnode.NewOperator().FromSlice(head).Head(), n)).ToSlice()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	head := []int{1, 2}
	n := 1
	expected := []int{1}
	actual, err := listnode.NewOperator().SetHead(removeNthFromEnd(listnode.NewOperator().FromSlice(head).Head(), n)).ToSlice()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
