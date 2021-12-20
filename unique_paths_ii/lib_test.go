package uniquepathsii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	exp := 2
	assert.Equal(t, exp, uniquePathsWithObstacles(obstacleGrid))
}

func TestSubmission1(t *testing.T) {
	obstacleGrid := [][]int{{0, 0}, {0, 1}}
	exp := 0
	assert.Equal(t, exp, uniquePathsWithObstacles(obstacleGrid))
}

func TestSubmission2(t *testing.T) {
	obstacleGrid := [][]int{{0, 1}}
	exp := 0
	assert.Equal(t, exp, uniquePathsWithObstacles(obstacleGrid))
}
