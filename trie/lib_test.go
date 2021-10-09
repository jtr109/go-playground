package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	word := "hello"
	obj := Constructor()
	obj.Insert(word)
}

func TestSearch(t *testing.T) {
	word := "hello"
	obj := Constructor()
	obj.Insert(word)
	assert.True(t, obj.Search(word))
	assert.False(t, obj.Search("he"))
}

func TestStartsWith(t *testing.T) {
	word := "hello"
	prefix1 := "he"
	prefix2 := "hex"
	obj := Constructor()
	obj.Insert(word)
	assert.True(t, obj.StartsWith(word))
	assert.True(t, obj.StartsWith(prefix1))
	assert.False(t, obj.StartsWith(prefix2))
}
