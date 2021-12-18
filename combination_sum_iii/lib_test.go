package combinationsumiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	k, n := 3, 7
	expected := [][]int{{1, 2, 4}}
	assert.Equal(t, expected, combinationSum3(k, n))
}
