package evaluatereversepolishnotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	expected := 9
	assert.Equal(t, expected, evalRPN(tokens))
}

func TestExample2(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	expected := 6
	assert.Equal(t, expected, evalRPN(tokens))
}

func TestExample3(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	expected := 22
	assert.Equal(t, expected, evalRPN(tokens))
}
