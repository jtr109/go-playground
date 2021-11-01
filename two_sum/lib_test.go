package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	assert.Equal(t, expected, twoSum(nums, target))
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	assert.Equal(t, expected, twoSum(nums, target))
}

func TestExample3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	assert.Equal(t, expected, twoSum(nums, target))
}
