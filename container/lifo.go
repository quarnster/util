// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

type (
	Lifo struct {
		Fifo
	}
)

func (li *Lifo) Less(i, j int) bool {
	return li.data[i].index > li.data[j].index
}
