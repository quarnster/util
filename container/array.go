// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container

import "github.com/quarnster/util"

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
	BasicArray struct {
		model []interface{}
	}
	ObservableArray struct {
		util.BasicObservable
		Array
	}
)

func (a *BasicArray) Insert(index int, data interface{}) error {
	if index < 0 {
		index = 0
	} else if index > len(a.model) {
		index = len(a.model)
	}

	nmodel := make([]interface{}, len(a.model)+1)
	copy(nmodel, a.model[:index])
	nmodel[index] = data
	copy(nmodel[index+1:], a.model[index:])
	return nil
}

func (a *BasicArray) Remove(i int) (olddata interface{}, err error) {
	if i < 0 {
		i = 0
	} else if i >= len(a.model)-1 {
		i = len(a.model) - 1
	}

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
