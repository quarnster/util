package container

import (
	"testing"
)

func TestFifo(t *testing.T) {
	items := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	pq := NewHeap(&Fifo{})
	for j := 0; j < 2; j++ {
		for _, t := range items {
			pq.Push(t)
		}
		for _, i := range items {
			v := pq.Pop().(int)
			if i != v {
				t.Errorf("%d !=  %d", i, v)
			}
		}
	}
}
