package intersectionoftwoarraysii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}
	assert.Equal(t, expected, intersect(nums1, nums2))
}

func TestExample2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{9, 4}
	assert.Equal(t, expected, intersect(nums1, nums2))
}
