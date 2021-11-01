package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(_t *testing.T) {
	s := "anagram"
	t := "nagaram"
	assert.True(_t, isAnagram(s, t))
}
