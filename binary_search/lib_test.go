package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	assert.Equal(t, expected, search(nums, target))
}

func TestExample2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	expected := -1
	assert.Equal(t, expected, search(nums, target))
}
