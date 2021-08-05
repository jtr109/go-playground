package validbraces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func singleTest(t *testing.T, str string, res bool) {
	assert.Equal(t, res, ValidBraces(str))
}

func TestValidBraces(t *testing.T) {
	singleTest(t, "(){}[]", true)
	singleTest(t, "([{}])", true)
	singleTest(t, "(}", false)
	singleTest(t, "[(])", false)
	singleTest(t, "[({)](]", false)
}
