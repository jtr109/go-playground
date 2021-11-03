package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := []byte("hello")
	expected := []byte("olleh")
	reverseString(s)
	assert.Equal(t, expected, s)
}
