package maximumnumberofballoons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 1, maxNumberOfBalloons("nlaebolko"))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 2, maxNumberOfBalloons("loonbalxballpoon"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, 0, maxNumberOfBalloons("leetcode"))
}
