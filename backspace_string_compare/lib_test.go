package backspacestringcompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "ab#c"
	t_ := "ab#c"
	assert.True(t, backspaceCompare(s, t_))
}

func TestExample2(t *testing.T) {
	s := "ab##"
	t_ := "c#d#"
	assert.True(t, backspaceCompare(s, t_))
}

func TestExample3(t *testing.T) {
	s := "a##c"
	t_ := "#a#c"
	assert.True(t, backspaceCompare(s, t_))
}

func TestExample4(t *testing.T) {
	s := "a#c"
	t_ := "b"
	assert.False(t, backspaceCompare(s, t_))
}
