package uniquebinarysearchtrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	n := 3
	exp := 5
	assert.Equal(t, exp, numTrees(n))
}

func TestExample2(t *testing.T) {
	n := 1
	exp := 1
	assert.Equal(t, exp, numTrees(n))
}

func TestExample3(t *testing.T) {
	n := 2
	exp := 2
	assert.Equal(t, exp, numTrees(n))
}
