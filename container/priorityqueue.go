// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

type PriorityQueue struct {
	Lifo
	Priority func(x interface{}) int
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.data = append(pq.data, data{pq.Priority(x), x})
}
