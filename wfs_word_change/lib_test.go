package wfswordchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGraph(t *testing.T) {
	arr := []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
	}
	graph := createGraph(arr)
	assert.Equal(t, "hot", graph[0].content)
}

func TestBuildGraph(t *testing.T) {
	arr := []string{
		"hit",
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}
	graph := createGraph(arr)
	buildGraph(graph)
	actual := []string{}
	for _, word := range graph[1].related {
		actual = append(actual, word.content)
	}
	exp := []string{"hit", "dot", "lot"} // hot
	assert.Equal(t, exp, actual)
}

func TestSearch(t *testing.T) {
	arr := []string{
		"hit",
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}
	graph := createGraph(arr)
	buildGraph(graph)
	begin := graph[0]
	end := graph[len(graph)-1]
	actual := search(begin, end)
	exp := 4
	assert.Equal(t, exp, actual)
}
