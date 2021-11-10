package slidingwindowmaximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expected := []int{3, 3, 5, 5, 6, 7}
	assert.Equal(t, expected, maxSlidingWindow(nums, k))
}

func TestExample2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	assert.Equal(t, expected, maxSlidingWindow(nums, k))
}

func TestExample3(t *testing.T) {
	nums := []int{1, -1}
	k := 1
	expected := []int{1, -1}
	assert.Equal(t, expected, maxSlidingWindow(nums, k))
}

func TestSubmission1(t *testing.T) {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	expected := []int{3, 3, 2, 5}
	assert.Equal(t, expected, maxSlidingWindow(nums, k))
}
