package maxpq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwim(t *testing.T) {
	pq := MaxPQ([]int{0, 2, 3, 1})
	pq.swim(2)
	expected := MaxPQ([]int{0, 3, 2, 1})
	assert.Equal(t, expected, pq)
}

func TestSwimDeeper(t *testing.T) {
	pq := MaxPQ([]int{0, 5, 3, 4, 6, 2, 1, 0})
	pq.swim(4)
	expected := MaxPQ([]int{0, 6, 5, 4, 3, 2, 1, 0})
	assert.Equal(t, expected, pq)
}

func TestSink(t *testing.T) {
	pq := MaxPQ([]int{0, 1, 2, 3})
	pq.sink(1)
	expected := MaxPQ([]int{0, 3, 2, 1})
	assert.Equal(t, expected, pq)
}

func TestSinkDeeper(t *testing.T) {
	pq := MaxPQ([]int{0, 1, 6, 5, 4, 3, 2})
	pq.sink(1)
	expected := MaxPQ([]int{0, 6, 4, 5, 1, 3, 2})
	assert.Equal(t, expected, pq)
}

func TestInsert(t *testing.T) {
	pq := MaxPQ([]int{0, 5, 4, 3, 2, 1})
	pq.Insert(6)
	expected := MaxPQ([]int{0, 6, 4, 5, 2, 1, 3})
	assert.Equal(t, expected, pq)
}

func TestPopMax(t *testing.T) {
	pq := MaxPQ([]int{0, 6, 4, 5, 2, 1, 3})
	res, ok := pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, res, 6)
	expected := MaxPQ([]int{0, 5, 4, 3, 2, 1})
	assert.Equal(t, expected, pq)
}

func TestMaxPQ(t *testing.T) {
	pq := NewMaxPQ()

	pq.Insert(4)
	pq.Insert(5)
	pq.Insert(1)

	res, ok := pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 5, res)

	pq.Insert(2)
	pq.Insert(6)
	pq.Insert(3)

	res, ok = pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 6, res)
	res, ok = pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 4, res)
	res, ok = pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 3, res)
	res, ok = pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 2, res)
	res, ok = pq.PopMax()
	assert.True(t, ok)
	assert.Equal(t, 1, res)

	res, ok = pq.PopMax()
	assert.False(t, ok)
	assert.Equal(t, 0, res)
}
