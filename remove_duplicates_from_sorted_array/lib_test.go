package removeduplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	assert.Equal(t, expected, removeDuplicates(nums))
}

func TestExample2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	assert.Equal(t, expected, removeDuplicates(nums))
}
