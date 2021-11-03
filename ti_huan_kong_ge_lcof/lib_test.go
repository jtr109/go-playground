package tihuankonggelcof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := []rune("We are happy.")
	expected := []rune("We%20are%20happy.")
	replaceSpace(&s)
	assert.Equal(t, expected, s)
}
