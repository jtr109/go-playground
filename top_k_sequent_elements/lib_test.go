package topksequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSequenceHeap(t *testing.T) {
	h := NewMaxSequenceHeap()
	e1 := NewElement(5, 1)
	e2 := NewElement(4, 2)
	e3 := NewElement(3, 3)
	e4 := NewElement(2, 4)
	e5 := NewElement(1, 5)
	h.Push(&e1)
	h.Push(&e5)
	h.Push(&e3)
	h.Push(&e2)
	h.Push(&e4)
	assert.Equal(t, &e5, h.Pop())
	assert.Equal(t, &e4, h.Pop())
	assert.Equal(t, &e3, h.Pop())
	assert.Equal(t, &e2, h.Pop())
	assert.Equal(t, &e1, h.Pop())
	assert.Nil(t, h.Pop())
}

func TestExample1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	assert.Equal(t, expected, topKFrequent(nums, k))
}

func TestExample2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	assert.Equal(t, expected, topKFrequent(nums, k))
}
