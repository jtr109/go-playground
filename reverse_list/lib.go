// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge-2021/636/week-1-september-1st-september-7th/3966/

package reverselist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode = nil
	for current := head; current != nil; current = current.Next {
		result = &ListNode{
			Val:  current.Val,
			Next: result,
		}
	}
	fmt.Println(result)
	return result
}
