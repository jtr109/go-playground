package foursumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	expected := 2
	assert.Equal(t, expected, fourSumCount(nums1, nums2, nums3, nums4))
}
