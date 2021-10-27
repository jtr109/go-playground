package reverselinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertListNodeToArray(head *ListNode) []int {
	current := head
	result := []int{}
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func convertArrayToListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &ListNode{
		Val: array[0],
	}
	current := head
	for i := 1; i < len(array); i++ {
		current.Next = &ListNode{
			Val: array[i],
		}
		current = current.Next
	}
	return head
}

func TestExample1(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	assert.Equal(t, expected, convertListNodeToArray(reverseList(convertArrayToListNode(head))))
}

func TestSubmission1(t *testing.T) {
	head := []int{}
	expected := []int{}
	assert.Equal(t, expected, convertListNodeToArray(reverseList(convertArrayToListNode(head))))
}
