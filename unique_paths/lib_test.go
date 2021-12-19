package uniquepaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	m, n := 3, 7
	exp := 28
	assert.Equal(t, exp, uniquePaths(m, n))
}
