// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

import (
	"testing"
)

func TestLifo(t *testing.T) {
	items := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	pq := NewHeap(&Lifo{})
	for j := 0; j < 2; j++ {
		for _, t := range items {
			pq.Push(t)
		}
		for i := range items {
			i = items[len(items)-1-i]
			v := pq.Pop().(int)
			if i != v {
				t.Errorf("%d !=  %d", i, v)
			}
		}
	}
}
