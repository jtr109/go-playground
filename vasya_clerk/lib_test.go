package vasyaclerk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTickets(t *testing.T) {
	assert.Equal(t, "YES", Tickets([]int{25, 25, 50}))
	assert.Equal(t, "NO", Tickets([]int{25, 100}))
}

func TestSubmission1(t *testing.T) {
	assert.Equal(t, "YES", Tickets([]int{25, 25, 50}))
	assert.Equal(t, "YES", Tickets([]int{25, 25, 50, 100}))
	assert.Equal(t, "NO", Tickets([]int{25, 100}))
	assert.Equal(t, "YES", Tickets([]int{25, 25, 25, 25, 25, 25, 25, 25, 25, 25}))
	assert.Equal(t, "NO", Tickets([]int{50, 50, 50, 50, 50, 50, 50, 50, 50, 50}))
	assert.Equal(t, "NO", Tickets([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100}))
	assert.Equal(t, "YES", Tickets([]int{25, 25, 25, 25, 50, 100, 50}))
	assert.Equal(t, "NO", Tickets([]int{50, 100, 100}))
	assert.Equal(t, "NO", Tickets([]int{25, 25, 100}))
}
