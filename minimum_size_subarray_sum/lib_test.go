package minimumsizesubarraysum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	expected := 2
	assert.Equal(t, expected, minSubArrayLen(target, nums))
}

func TestExample2(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	expected := 1
	assert.Equal(t, expected, minSubArrayLen(target, nums))
}
