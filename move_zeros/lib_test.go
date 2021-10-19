package movezeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func TestExample2(t *testing.T) {
	nums := []int{0}
	expected := []int{0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}
