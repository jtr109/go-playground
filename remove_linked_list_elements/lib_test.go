package removelinkedlistelements

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
	head := convertArrayToListNode([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	removeElements(head, val)
	result := convertListNodeToArray(head)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}
