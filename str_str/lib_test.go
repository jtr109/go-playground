package strstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	expected := 2
	assert.Equal(t, expected, strStr(haystack, needle))
}

func TestExample2(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	expected := -1
	assert.Equal(t, expected, strStr(haystack, needle))
}

func TestExample3(t *testing.T) {
	haystack := ""
	needle := ""
	expected := 0
	assert.Equal(t, expected, strStr(haystack, needle))
}

func TestSubmission1(t *testing.T) {
	haystack := "aaa"
	needle := "aaaa"
	expected := -1
	assert.Equal(t, expected, strStr(haystack, needle))
}

func TestBuildNext(t *testing.T) {
	needle := "aabaaf"
	expected := []int{-1, 0, -1, 0, 1, -1}
	assert.Equal(t, expected, buildNext(needle))
}
