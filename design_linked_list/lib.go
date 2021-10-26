// https://leetcode.com/problems/design-linked-list/

package designlinkedlist

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val: -1,
	}
}

func (this *MyLinkedList) Get(index int) int {
	current := this
	for i := -1; i < index; i++ {
		if current.Next == nil {
			// out of range
			return -1
		}
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := MyLinkedList{
		Val:  val,
		Next: this.Next,
	}
	this.Next = &node
}

func (this *MyLinkedList) AddAtTail(val int) {
	current := this
	for current.Next != nil {
		current = current.Next
	}
	node := MyLinkedList{
		Val: val,
	}
	current.Next = &node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	current := this
	for i := -1; i < index-1; i++ {
		if current.Next == nil {
			// out of range
			return
		}
		current = current.Next
	}
	current.Next = &MyLinkedList{
		Val:  val,
		Next: current.Next,
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	current := this
	for i := -1; i < index-1; i++ {
		if current.Next == nil {
			// out of range
			return
		}
		current = current.Next
	}
	if current.Next == nil {
		return
	}
	current.Next = current.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
