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
	head := convertArrayToListNode([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, convertListNodeToArray(removeElements(head, val)))
}

func TestExample2(t *testing.T) {
	head := convertArrayToListNode([]int{})
	val := 1
	expected := []int{}
	assert.Equal(t, expected, convertListNodeToArray(removeElements(head, val)))
}

func TestExample3(t *testing.T) {
	head := convertArrayToListNode([]int{7, 7, 7, 7})
	val := 7
	expected := []int{}
	assert.Equal(t, expected, convertListNodeToArray(removeElements(head, val)))
}
