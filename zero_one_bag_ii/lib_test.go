package zeroonebagii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	targetWeight := 4
	exp := 35
	assert.Equal(t, exp, getMaxValue(weight, value, targetWeight))
}
