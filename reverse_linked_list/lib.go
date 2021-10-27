package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var result *ListNode = nil
	for current := head; current != nil; current = current.Next {
		result = &ListNode{
			Val:  current.Val,
			Next: result,
		}
	}
	return result
}

// time complexity: O(N)
// space complexity: O(N)

func reverseList2(head *ListNode) *ListNode {
	newHead, _ := reverseLinkedList(head)
	return newHead
}

func reverseLinkedList(head *ListNode) (newHead *ListNode, newTail *ListNode) {
	if head == nil || head.Next == nil { // original tail
		return head, head
	}
	newHead, nextNewTail := reverseLinkedList(head.Next)
	newTail = head
	newTail.Next = nil
	nextNewTail.Next = newTail
	return
}

// time complexity: O(N)
// space complexity: O(1)
// 但是这种方式会产生递归调用栈，也会消耗空间

// 双指针模式
// time complexity: O(N)
// space complexity: O(1)
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode = nil
	current := head
	for current != nil {
		temp := current
		current = current.Next
		temp.Next = previous
		previous = temp
	}
	return previous
}
