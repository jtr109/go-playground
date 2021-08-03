package countthedigit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
	assert.Equal(t, []int{1, 2}, digits(21))
	assert.Equal(t, []int{0}, digits(0))
	assert.Equal(t, []int{0, 0, 1}, digits(100))
}
