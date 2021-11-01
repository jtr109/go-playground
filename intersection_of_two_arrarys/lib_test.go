package intersectionoftwoarrarys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2}
	assert.Equal(t, expected, intersection(nums1, nums2))
}
