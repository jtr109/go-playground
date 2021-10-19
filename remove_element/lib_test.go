package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2
	assert.Equal(t, expected, removeElement(nums, val))
}

func TestExample2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expected := 5
	assert.Equal(t, expected, removeElement(nums, val))
}
