// https://leetcode.com/problems/sliding-window-maximum/

package slidingwindowmaximum

type Deque struct {
	elements []int
}

func (d *Deque) Empty() bool {
	return len(d.elements) == 0
}

func (d *Deque) Right() (int, bool) {
	if d.Empty() {
		return 0, false
	}
	return d.elements[len(d.elements)-1], true
}

func (d *Deque) PopRight() (int, bool) {
	if d.Empty() {
		return 0, false
	}
	n := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return n, true
}

func (d *Deque) PushRight(n int) {
	for {
		r, ok := d.Right()
		if !ok {
			break
		}
		if n <= r {
			break
		}
		d.PopRight()
	}
	d.elements = append(d.elements, n)
}

func (d *Deque) Left() (int, bool) {
	if d.Empty() {
		return 0, false
	}
	return d.elements[0], true
}

func (d *Deque) PopLeft() (int, bool) {
	if d.Empty() {
		return 0, false
	}
	n := d.elements[0]
	d.elements = d.elements[1:]
	return n, true
}

func NewDeque() *Deque {
	return &Deque{}
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := NewDeque()
	result := []int{}
	for i := 0; i < k; i++ {
		deque.PushRight(nums[i])
	}
	n, _ := deque.Left()
	result = append(result, n)
	for i := k; i < len(nums); i++ {
		left, _ := deque.Left()
		if nums[i-k] == left {
			deque.PopLeft()
		}
		deque.PushRight(nums[i])
		left, _ = deque.Left()
		result = append(result, left)
	}
	return result
}
