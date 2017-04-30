package algorithm

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	h := &MinHeap{}
	h.values = []int{1, 2, 3}
	h.Swap(0, 2)
	if h.values[0] != 3 || h.values[1] != 2 || h.values[2] != 1 {
		t.Error("failed swap")
	}
}

func TestPop(t *testing.T) {
	h := &MinHeap{}
	h.values = []int{1}
	if h.Pop() != 1 {
		t.Error("failed pop")
	}

	h.values = []int{1, 2, 3}
	h.Pop()
	if h.values[0] != 2 {
		fmt.Println(h.values)
		t.Error("failed pop")
	}

	h.values = []int{1, 3, 2}
	h.Pop()
	if h.values[0] != 2 {
		t.Error("failed pop")
	}
}

func TestPush(t *testing.T) {
	h := &MinHeap{}
	h.values = []int{}
	h.Push(1)
	if h.values[0] != 1 {
		t.Error("failed push")
	}

	h.values = []int{1}
	h.Push(2)
	if h.values[0] != 1 {
		t.Error("failed push")
	}

	h.values = []int{2}
	h.Push(1)
	if h.values[0] != 1 {
		t.Error("failed push")
	}
}

func TestSiftDown(t *testing.T) {
	h := &MinHeap{}

	h.values = []int{1}
	h.siftDown()
	if h.values[0] != 1 {
		t.Error("failed sift down")
	}
	h.values = []int{2, 1}
	h.siftDown()
	if h.values[0] != 1 {
		t.Error("failed sift down")
	}

	h.values = []int{1, 2, 3}
	h.siftDown()
	if h.values[0] != 1 {
		t.Error("failed sift down")
	}

	h.values = []int{2, 1, 3}
	h.siftDown()
	if h.values[0] != 1 {
		t.Error("failed sift down")
	}

	h.values = []int{3, 2, 1}
	h.siftDown()
	if h.values[0] != 1 {
		t.Error("failed sift down")
	}
}

func TestSiftUp(t *testing.T) {
	h := &MinHeap{}

	h.values = []int{1}
	h.siftUp(0)
	if h.values[0] != 1 {
		t.Error("failed sift up")
	}

	h.values = []int{2, 1}
	h.siftUp(1)
	if h.values[0] != 1 {
		t.Error("failed sift up")
	}

	h.values = []int{1, 2, 3}
	h.siftUp(2)
	if h.values[0] != 1 {
		t.Error("failed sift up")
	}

	h.values = []int{2, 1, 3}
	h.siftUp(1)
	if h.values[0] != 1 {
		t.Error("failed sift up")
	}

	h.values = []int{3, 2, 1}
	h.siftUp(2)
	if h.values[0] != 1 {
		t.Error("failed sift up")
	}
}
