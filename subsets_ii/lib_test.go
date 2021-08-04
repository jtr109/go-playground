package subsetsii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetWithDup(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, subsetsWithDup([]int{1, 2, 3}))
	assert.Equal(t, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}, subsetsWithDup([]int{1, 2, 2}))
}

func TestSubsetWithDupSubmission1(t *testing.T) {
	expected := [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}}
	sort.Slice(expected, func(i, j int) bool {
		if len(expected[i]) != len(expected[j]) {
			return len(expected[i]) < len(expected[j])
		}
		for k := range expected[i] {
			if expected[i][k] < expected[j][k] {
				return true
			}
		}
		return false
	})
	actual := subsetsWithDup([]int{2, 1, 2, 1, 3})
	sort.Slice(actual, func(i, j int) bool {
		if len(actual[i]) != len(actual[j]) {
			return len(actual[i]) < len(actual[j])
		}
		for k := range actual[i] {
			if actual[i][k] < actual[j][k] {
				return true
			}
		}
		return false
	})
	assert.Equal(t, expected, actual)
}
