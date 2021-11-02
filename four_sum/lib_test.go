package foursum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	expected := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	assert.Equal(t, expected, fourSum(nums, target))
}

func TestExample2(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	target := 8
	expected := [][]int{{2, 2, 2, 2}}
	assert.Equal(t, expected, fourSum(nums, target))
}
