package happynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	n := 19
	assert.True(t, isHappy(n))
}

func TestExample2(t *testing.T) {
	n := 2
	assert.False(t, isHappy(n))
}

func TestSubmission1(t *testing.T) {
	n := 7
	assert.True(t, isHappy(n))
}
