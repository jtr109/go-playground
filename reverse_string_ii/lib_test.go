package reversestringii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "bacdfeg"
	assert.Equal(t, expected, reverseStr(s, k))
}
