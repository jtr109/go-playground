package mincostclimbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	cost := []int{10, 15, 20}
	expected := 15
	assert.Equal(t, expected, minCostClimbingStairs(cost))
}
