// https://leetcode.com/problems/top-k-frequent-elements/

package topksequentelements

type Element struct {
	value    int
	sequence int
}

func NewElement(value int, sequence int) Element {
	return Element{value: value, sequence: sequence}
}

type MaxSequenceHeap struct {
	heap []*Element
}

func (h *MaxSequenceHeap) swim(i int) {
	for ; i >= 0 && h.heap[i].sequence > h.heap[i/2].sequence; i /= 2 {
		h.heap[i], h.heap[i/2] = h.heap[i/2], h.heap[i]
	}
}

func (h *MaxSequenceHeap) sink(i int) {
	for {
		j := i * 2
		if j >= len(h.heap) {
			break
		}
		if j+1 < len(h.heap) { // element h.heap[j+1] exists
			if h.heap[j+1].sequence > h.heap[j].sequence {
				j++
			}
		}
		// j is the larger element of the two children
		if h.heap[i].sequence >= h.heap[j].sequence {
			break
		}
		// the element in index j have more sequence than the one at index i
		h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
		i = j // next i
	}
}

func (h *MaxSequenceHeap) isEmpty() bool {
	return len(h.heap) == 0
}

func (h *MaxSequenceHeap) Push(e *Element) {
	h.heap = append(h.heap, e)
	h.swim(len(h.heap) - 1)
}

func (h *MaxSequenceHeap) Pop() *Element {
	if h.isEmpty() {
		return nil
	}
	result := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.sink(0)
	return result
}

func NewMaxSequenceHeap() MaxSequenceHeap {
	return MaxSequenceHeap{}
}

func topKFrequent(nums []int, k int) []int {
	sequences := make(map[int]int)
	for _, n := range nums {
		sequences[n]++
	}
	h := NewMaxSequenceHeap()
	for n, seq := range sequences {
		h.Push(&Element{
			value:    n,
			sequence: seq,
		})
	}
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, h.Pop().value)
	}
	return result
}
