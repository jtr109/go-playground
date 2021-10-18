package validperfectsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	num := 16
	assert.True(t, isPerfectSquare(num))
}

func TestExample2(t *testing.T) {
	num := 14
	assert.False(t, isPerfectSquare(num))
}
