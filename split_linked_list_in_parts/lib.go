package splitlinkedlistinparts

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) (res []*ListNode) {
	current := head
	for _, length := range lengthList(head, k) {
		res = append(res, current)
		current = cut(current, length)
	}
	return
}

func count(ln *ListNode) (count int) {
	for current := ln; current.Next != nil; current = current.Next {
		count++
	}
	count++
	return
}

func lengthList(ln *ListNode, k int) (res []int) {
	count := count(ln)
	shorterLength := count / k
	largerCount := count % k // number of elements have larger length
	for i := 0; i < k; i++ {
		if i < largerCount {
			res = append(res, shorterLength+1)
		} else {
			res = append(res, shorterLength)
		}
	}
	return
}

func cut(ln *ListNode, length int) (next *ListNode) {
	current := ln
	if current == nil {
		return nil
	}
	for i := 0; i < length; i++ {
		if current == nil {
			return nil
		}
		if i == length-1 {
			// the current one is the last element
			break
		}
		current = current.Next
	}
	next = current.Next
	current.Next = nil
	return next
}
