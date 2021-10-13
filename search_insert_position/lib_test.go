package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	assert.Equal(t, expected, searchInsert(nums, target))
}

func TestExample2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	expected := 1
	assert.Equal(t, expected, searchInsert(nums, target))
}

func TestExample3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	expected := 4
	assert.Equal(t, expected, searchInsert(nums, target))
}

func TestExample4(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 0
	expected := 0
	assert.Equal(t, expected, searchInsert(nums, target))
}

func TestExample5(t *testing.T) {
	nums := []int{1}
	target := 0
	expected := 0
	assert.Equal(t, expected, searchInsert(nums, target))
}

func TestSubmission1(t *testing.T) {
	nums := []int{1}
	target := 1
	expected := 0
	assert.Equal(t, expected, searchInsert(nums, target))
}
