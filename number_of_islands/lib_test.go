package numberofislands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	grid := [][]byte{
		{byte(1), byte(1), byte(1), byte(1), byte(0)},
		{byte(1), byte(1), byte(0), byte(1), byte(0)},
		{byte(1), byte(1), byte(0), byte(0), byte(0)},
		{byte(0), byte(0), byte(0), byte(0), byte(0)},
	}
	expected := 1
	assert.Equal(t, expected, numIslands(grid))
}

func TestSubmission1(t *testing.T) {
	grid := [][]byte{
		{byte(1), byte(1), byte(0), byte(0), byte(0)},
		{byte(1), byte(1), byte(0), byte(0), byte(0)},
		{byte(0), byte(0), byte(1), byte(0), byte(0)},
		{byte(0), byte(0), byte(0), byte(1), byte(1)},
	}
	expected := 3
	assert.Equal(t, expected, numIslands(grid))
}
