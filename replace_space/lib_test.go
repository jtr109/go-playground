package replacespace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpace(t *testing.T) {
	s1 := "Good morning, sir!"
	expected1 := "Good%20morning,%20sir!"
	assert.Equal(t, expected1, replaceSpace(s1))
}
