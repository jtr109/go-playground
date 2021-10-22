package minimumwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t_ *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"
	expected := "BANC"
	assert.Equal(t_, expected, minWindow(s, t))
}

func TestExample2(t_ *testing.T) {
	s := "a"
	t := "a"
	expected := "a"
	assert.Equal(t_, expected, minWindow(s, t))
}

func TestExample3(t_ *testing.T) {
	s := "a"
	t := "aa"
	expected := ""
	assert.Equal(t_, expected, minWindow(s, t))
}

func TestSubmission1(t_ *testing.T) {
	s := "a"
	t := "b"
	expected := ""
	assert.Equal(t_, expected, minWindow(s, t))
}
