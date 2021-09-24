package nthtribonaccinumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 4, tribonacci(4))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 1389537, tribonacci(25))
}
