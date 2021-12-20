package integerbreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	n := 2
	exp := 1
	assert.Equal(t, exp, integerBreak(n))
}

func TestExample2(t *testing.T) {
	n := 10
	exp := 36
	assert.Equal(t, exp, integerBreak(n))
}
