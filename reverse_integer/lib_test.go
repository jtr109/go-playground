package reverseinteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	a := 123
	exp := 321
	assert.Equal(t, exp, reverse(a))
}

func TestReverse1(t *testing.T) {
	assert.Equal(t, 0, reverse(0))
	assert.Equal(t, 1, reverse(10))
	assert.Equal(t, 102, reverse(201))
}

func TestNegative(t *testing.T) {
	assert.Equal(t, -1, reverse(-1))
	assert.Equal(t, -12, reverse(-21))
	assert.Equal(t, -1, reverse(-10))
	assert.Equal(t, 0, reverse(7223372036854775899))  // 9223372036854775807
	assert.Equal(t, 0, reverse(-8023372036854775899)) // -9223372036854775808
}
