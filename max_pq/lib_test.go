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
