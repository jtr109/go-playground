package shiftingletters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharShifts(t *testing.T) {
	shifts := []int{3, 5, 9}
	expected := []int{3 + 5 + 9, 5 + 9, 9}
	assert.Equal(t, expected, getCharShifts(shifts))
}

func TestShiftMultipleTimes(t *testing.T) {
	assert.Equal(t, 'b', shiftMultipleTimes('a', 1))
	assert.Equal(t, 'u', shiftMultipleTimes('t', 1))
	assert.Equal(t, 'a', shiftMultipleTimes('z', 1))
	assert.Equal(t, 'r', shiftMultipleTimes('a', 3+5+9))
	assert.Equal(t, 'p', shiftMultipleTimes('b', 5+9))
	assert.Equal(t, 'l', shiftMultipleTimes('c', 9))
}

func TestExample1(t *testing.T) {
	s := "abc"
	shifts := []int{3, 5, 9}
	expected := "rpl"
	assert.Equal(t, expected, shiftingLetters(s, shifts))
}

func TestExample2(t *testing.T) {
	s := "aaa"
	shifts := []int{1, 2, 3}
	expected := "gfd"
	assert.Equal(t, expected, shiftingLetters(s, shifts))
}
