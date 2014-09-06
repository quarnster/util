// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"github.com/quarnster/util"
	"sort"
)

var (
	ErrNotInt           = fmt.Errorf("Attempting to insert a non-int type")
	ErrIndexOOB         = fmt.Errorf("Index is out of bounds")
	ErrNotManipulatable = fmt.Errorf("Filtered arrays are not directly manipulatable")
)

type (
	RemovedData struct {
		Index int
		Data  interface{}
	}

	InsertedData struct {
		Index int
	}

	Array interface {
		Insert(index int, data interface{}) error
		Remove(i int) (olddata interface{}, err error)
		Get(index int) interface{}
		Len() int
	}
	IntArray struct {
		BasicArray
	}
	BasicArray struct {
		model []interface{}
	}
	BoundsCheckingArray struct {
		Array
	}
	ObservableArray struct {
		util.BasicObservable
		Array
	}
	Acceptable    func(data interface{}) bool
	filteredArray struct {
		indices IntArray
		accept  Acceptable
		Array
	}
)

func NewFilteredArray(inner Array, accept Acceptable) Array {
	if _, ok := inner.(util.Observable); !ok {
		inner = &ObservableArray{Array: inner}
	}
	fa := filteredArray{accept: accept, Array: inner}
	inner.(util.Observable).AddObserver(&fa)
	for i := 0; i < inner.Len(); i++ {
		if !accept(inner.Get(i)) {
			continue
		}
		fa.indices.Insert(fa.indices.Len(), i)
	}
	return &fa
}

func (b *BoundsCheckingArray) Insert(index int, data interface{}) error {
	if index < 0 || index > b.Len() {
		return ErrIndexOOB
	}
	return b.Array.Insert(index, data)
}

func (b *BoundsCheckingArray) Remove(index int) (interface{}, error) {
	if index < 0 || index >= b.Len() {
		return nil, ErrIndexOOB
	}
	return b.Array.Remove(index)
}

func (b *BoundsCheckingArray) Get(index int) interface{} {
	if index < 0 || index >= b.Len() {
		return nil
	}
	return b.Array.Get(index)
}

func (i *IntArray) Insert(index int, data interface{}) error {
	_, ok := data.(int)
	if !ok {
		return ErrNotInt
	}
	return i.BasicArray.Insert(index, data)
}

func (a *BasicArray) Insert(index int, data interface{}) error {
	nmodel := make([]interface{}, len(a.model)+1)
	copy(nmodel, a.model[:index])
	nmodel[index] = data
	copy(nmodel[index+1:], a.model[index:])
	return nil
}

func (a *BasicArray) Remove(i int) (olddata interface{}, err error) {
	olddata = a.model[i]
	copy(a.model[i:], a.model[i+1:])
	return olddata, nil
}

func (a *BasicArray) Get(index int) interface{} {
	return a.model[index]
}

func (a *BasicArray) Len() int {
	return len(a.model)
}

func (a *ObservableArray) Insert(index int, data interface{}) error {
	if err := a.Array.Insert(index, data); err != nil {
		return err
	}
	a.NotifyObservers(InsertedData{index})
	return nil
}

func (a *ObservableArray) Remove(i int) (olddata interface{}, err error) {
	if olddata, err = a.Array.Remove(i); err != nil {
		return
	}
	a.NotifyObservers(RemovedData{i, olddata})
	return
}

func (fa *filteredArray) Changed(data interface{}) {
	switch d := data.(type) {
	case RemovedData:
		for i, k := range fa.indices.model {
			if k == d.Index {
				fa.indices.Remove(i)
				break
			}
		}
	case InsertedData:
		data := fa.Get(d.Index)
		if !fa.accept(data) {
			return
		}
		idx := sort.Search(fa.indices.Len(), func(i int) bool {
			return fa.Get(i).(int) < d.Index
		})
		fa.indices.Insert(idx+1, d.Index)
	}
}

func (fa *filteredArray) Insert(index int, data interface{}) error {
	return ErrNotManipulatable
}

func (fa *filteredArray) Remove(index int) (interface{}, error) {
	return nil, ErrNotManipulatable
}

func (fa *filteredArray) Get(index int) interface{} {
	return fa.Array.Get(fa.indices.Get(index).(int))
}
func (fa *filteredArray) Len() int {
	return fa.indices.Len()
}
