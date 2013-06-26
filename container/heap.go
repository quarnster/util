package container

import (
	"container/heap"
)

type (
	Heap interface {
		Push(interface{})
		Pop() interface{}
		Len() int
	}
	h struct {
		hi heap.Interface
	}
)

func NewHeap(hi heap.Interface) Heap {
	heap.Init(hi)
	return &h{hi}
}

func (h *h) Len() int {
	return h.hi.Len()
}

func (h *h) Push(x interface{}) {
	heap.Push(h.hi, x)
}
func (h *h) Pop() interface{} {
	return heap.Pop(h.hi)
}
