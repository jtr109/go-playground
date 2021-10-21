package fruitintobackets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	fruits := []int{1, 2, 1}
	expected := 3
	assert.Equal(t, expected, totalFruit(fruits))
}

func TestExample2(t *testing.T) {
	fruits := []int{0, 1, 2, 2}
	expected := 3
	assert.Equal(t, expected, totalFruit(fruits))
}

func TestExample3(t *testing.T) {
	fruits := []int{1, 2, 3, 2, 2}
	expected := 4
	assert.Equal(t, expected, totalFruit(fruits))
}

func TestExample4(t *testing.T) {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	expected := 5
	assert.Equal(t, expected, totalFruit(fruits))
}
