package spiralmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	assert.Equal(t, expected, spiralOrder(matrix))
}

func TestExample2(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	assert.Equal(t, expected, spiralOrder(matrix))
}
