package repeatedsubstringpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildNext(t *testing.T) {
	needle := "aabaaf"
	expected := []int{-1, 0, -1, 0, 1, -1}
	assert.Equal(t, expected, buildNext(needle))
}

func TestExample1(t *testing.T) {
	s := "abab"
	assert.True(t, repeatedSubstringPattern(s))
}

func TestExample2(t *testing.T) {
	s := "aba"
	assert.False(t, repeatedSubstringPattern(s))
}

func TestExample3(t *testing.T) {
	s := "abcabcabcabc"
	assert.True(t, repeatedSubstringPattern(s))
}

func TestSubmission1(t *testing.T) {
	s := "abac"
	assert.False(t, repeatedSubstringPattern(s))
}
