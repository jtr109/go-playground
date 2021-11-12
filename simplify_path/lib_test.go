package simplifypath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	input := "/home/"
	expected := "/home"
	actual := simplifyPath(input)
	assert.Equal(t, expected, actual)
}

func TestExample2(t *testing.T) {
	input := "/../"
	expected := "/"
	actual := simplifyPath(input)
	assert.Equal(t, expected, actual)
}

func TestExample3(t *testing.T) {
	input := "/home//foo/"
	expected := "/home/foo"
	actual := simplifyPath(input)
	assert.Equal(t, expected, actual)
}
