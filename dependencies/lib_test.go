package dependencies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGraph(t *testing.T) {

	dependencies := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
	}
	operations := []*Operation{}
	for i := 0; i < 4; i++ {
		op := &Operation{
			ID:           i,
			dependencies: []*Operation{},
		}
		operations = append(operations, op)
	}
	createGraph(operations, dependencies)
	assert.Equal(t, []*Operation{operations[1], operations[2]}, operations[0].dependencies)
}

func TestDFS(t *testing.T) {
	dependencies := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
	}
	operations := []*Operation{}
	for i := 0; i < 4; i++ {
		op := &Operation{
			ID:           i,
			dependencies: []*Operation{},
		}
		operations = append(operations, op)
	}
	createGraph(operations, dependencies)

	closeList = []int{}
	for _, op := range operations {
		dfs(op)
	}
	exp := []int{3, 1, 2, 0}
	assert.Equal(t, exp, closeList)
}
