package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "()"
	assert.True(t, isValid(s))
}

func TestExample2(t *testing.T) {
	s := "()[]{}"
	assert.True(t, isValid(s))
}

func TestExample3(t *testing.T) {
	s := "(]"
	assert.False(t, isValid(s))
}

func TestSubmission1(t *testing.T) {
	s := "]"
	assert.False(t, isValid(s))
}
