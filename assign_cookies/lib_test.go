package assigncookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	expected := 1
	assert.Equal(t, expected, findContentChildren(g, s))
}

func TestSubmission1(t *testing.T) {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	expected := 2
	assert.Equal(t, expected, findContentChildren(g, s))
}
