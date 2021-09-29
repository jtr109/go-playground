package splitlinkedlistinparts

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) (res []*ListNode) {
	current := head
	for _, length := range lengthList(head, k) {
		res = append(res, current)
		current = move(current, length)
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

func move(ln *ListNode, step int) *ListNode {
	current := ln
	for i := 0; i < step; i++ {
		current = current.Next
	}
	return current
}
