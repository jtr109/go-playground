package sameaftersorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameAfterSorted(t *testing.T) {
	s1 := "abcabd"
	s2 := "abcbda"
	s3 := "abcbdc"
	assert.True(t, sameAfterSorted(s1, s2))
	assert.False(t, sameAfterSorted(s1, s3))
}
