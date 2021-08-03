package string_end_with

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, true, solution("", ""))
	assert.Equal(t, true, solution("abc", "c"))
}
