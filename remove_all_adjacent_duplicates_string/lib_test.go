package removealladjacentduplicatesstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "abbaca"
	expected := "ca"
	assert.Equal(t, expected, removeDuplicates(s))
}

func TestExample2(t *testing.T) {
	s := "azxxzy"
	expected := "ay"
	assert.Equal(t, expected, removeDuplicates(s))
}
