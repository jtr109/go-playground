package designlinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	assert.Equal(t, 2, obj.Get(1))
	obj.DeleteAtIndex(1)
	assert.Equal(t, 3, obj.Get(1))
}

func TestSubmission1(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	obj.Get(5)
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)
}
