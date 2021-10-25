package spiralmatrixii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeMatrix(t *testing.T) {
	n := 3
	expected := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	assert.Equal(t, expected, initializeMatrix(n))
}

func TestExample1(t *testing.T) {
	n := 3
	expected := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	assert.Equal(t, expected, generateMatrix(n))
}
