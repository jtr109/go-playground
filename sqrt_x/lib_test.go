package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	x := 4
	expected := 2
	assert.Equal(t, expected, mySqrt(x))
}

func TestExample2(t *testing.T) {
	x := 8
	expected := 2
	assert.Equal(t, expected, mySqrt(x))
}
