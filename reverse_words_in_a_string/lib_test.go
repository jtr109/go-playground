package reversewordsinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "the sky is blue"
	expected := "blue is sky the"
	assert.Equal(t, expected, reverseWords(s))
}

func TestExample2(t *testing.T) {
	s := "  hello world  "
	expected := "world hello"
	assert.Equal(t, expected, reverseWords(s))
}
