package lettercombinationsofaphonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, expected, letterCombinations(digits))
}
