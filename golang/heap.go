package algorithm

import "sort"

type heap interface {
	Pop() int
	Push(int)
	sort.Interface
}

type MinHeap struct {
	values []int
}

func (h *MinHeap) Len() int {
	return len(h.values)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.values[i] < h.values[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *MinHeap) Pop() int {
	l := len(h.values)
	h.Swap(0, l-1)
	val := h.values[l-1]
	h.values = h.values[:l-1]
	h.siftDown()
	return val
}

func (h *MinHeap) siftDown() {
	l := len(h.values) - 1
	i := 0
	for {
		c1 := i*2 + 1
		c2 := c1 + 1

		if c1 > l {
			break
		}

		c := c1
		if c2 <= l && h.values[c1] > h.values[c2] {
			c = c2
		}

		if h.values[c] > h.values[i] {
			break
		}
		h.Swap(i, c)
		i = c
	}
}

func (h *MinHeap) Push(v int) {
	h.values = append(h.values, v)
	h.siftUp(len(h.values) - 1)
}

func (h *MinHeap) siftUp(i int) {
	p := i / 2
	if h.values[p] > h.values[i] {
		h.Swap(p, i)
		h.siftUp(p)
	}
}
