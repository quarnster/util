// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

import (
	"testing"
)

func TestPQ(t *testing.T) {
	type test struct {
		value    string
		priority int
	}
	items := []test{
		{"a", 99},
		{"b", 98},
		{"c", 97},
		{"d", 100},
		{"e", 101},
		{"f", 102},
	}
	pq := NewHeap(&PriorityQueue{Priority: func(x interface{}) int {
		return x.(test).priority
	}})
	for _, t := range items {
		pq.Push(t)
	}
	min := 10000
	for pq.Len() != 0 {
		v := pq.Pop().(test).priority
		t.Log(v)
		if min < v {
			t.Errorf("%d < %d", min, v)
		} else {
			min = v
		}
	}
}
