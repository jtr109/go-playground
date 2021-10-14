package findfirstandlastpositionofelementinsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchLeft1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := 3
	assert.Equal(t, expected, searchLeft(nums, target))
}

func TestSearchLeft2(t *testing.T) {
	nums := []int{5, 7, 8, 8, 8, 10}
	target := 8
	expected := 2
	assert.Equal(t, expected, searchLeft(nums, target))
}

func TestSearchLeft3(t *testing.T) {
	nums := []int{5, 8, 8, 8, 8, 10}
	target := 8
	expected := 1
	assert.Equal(t, expected, searchLeft(nums, target))
}

func TestSearchLeft4(t *testing.T) {
	nums := []int{8, 8, 8, 8, 8, 10}
	target := 8
	expected := 0
	assert.Equal(t, expected, searchLeft(nums, target))
}

func TestSearchLeftDoesNotExist(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := -1
	assert.Equal(t, expected, searchLeft(nums, target))
}

func TestSearchRight1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := 4
	assert.Equal(t, expected, searchRight(nums, target))
}

func TestSearchRightDoesNotExist(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := -1
	assert.Equal(t, expected, searchRight(nums, target))
}

func TestExample1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	assert.Equal(t, expected, searchRange(nums, target))
}

func TestExample2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1}
	assert.Equal(t, expected, searchRange(nums, target))

}

func TestExample3(t *testing.T) {
	nums := []int{}
	target := 0
	expected := []int{-1, -1}
	assert.Equal(t, expected, searchRange(nums, target))
}
