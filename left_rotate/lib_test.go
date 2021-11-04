package leftrotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "cdefgab"
	assert.Equal(t, expected, leftRotate(s, k))
}

func TestExample2(t *testing.T) {
	s := "lrloseumgh"
	k := 6
	expected := "umghlrlose"
	assert.Equal(t, expected, leftRotate(s, k))
}
