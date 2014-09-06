// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"reflect"
)

type (
	Getfunc func() interface{}
	// Defines an interface for Getable types
	Getable interface {
		// Should return the current value of this type
		Get() interface{}
	}
	// Defines an interface for Setable types
	Setable interface {
		// Should set the value of the current type to
		// the provided interface value.
		Set(v interface{})
	}
	// Classic observer callback
	Observer interface {
		Changed(data interface{})
	}

	// A type that is observable keeps its own list of
	// active observers which should be notified upon
	// changes being made to this type
	Observable interface {
		AddObserver(Observer)
		RemoveObserver(Observer)
		NotifyObservers(data interface{})
	}

	// BasicObservable implements the Observable interface,
	// and is suitable to use as an embedded struct.
	BasicObservable struct {
		observers []Observer
	}

	// The "poll" type deals with observed values that do
	// not satisfy the Observable interface.
	//
	// For these values we have to actively check their value
	// and compare it to a stored value to detect whether it
	// has changed or not. Aka polling.
	poll struct {
		BasicObservable
		current interface{}
		get     Getfunc
	}

	// The Observatory type deals with connecting source and
	// target values (via Connect) and manually polling for
	// source value changes if needed.
	//
	// Users will have to call the Update function as appropriate
	// for polling and comparing the non-observable source values
	// that are hooked up with the help of this Observatory.
	Observatory struct {
		pollers map[interface{}]poll
	}

	funcObserver struct {
		obs func()
	}
)

func (o funcObserver) Changed(data interface{}) {
	o.obs()
}
func (o *BasicObservable) AddObserver(obs Observer) {
	o.observers = append(o.observers, obs)
}

func (o *BasicObservable) NotifyObservers(data interface{}) {
	for _, obs := range o.observers {
		obs.Changed(data)
	}
}

// Creates a function that returns the current value of the provided source parameter.
func (o *Observatory) CreateGetfunc(source interface{}) (get Getfunc) {
	if sg, ok := source.(Getable); ok {
		get = sg.Get
	} else if t := reflect.ValueOf(source); t.Kind() == reflect.Ptr {
		get = func() interface{} {
			return t.Elem().Interface()
		}
	} else {
		panic(fmt.Errorf("Can't deal with that non-getable source type: %s", reflect.TypeOf(source)))
	}
	return
}

// Observes a source value for changes, invoking the provided Observer upon changes.
// If the source is not satisfying the Observable interface, the source will be added
// to this Observatory's local polling list where it will use "get" and compare the
// returned value with a stored value copy.
func (o *Observatory) Observe(source interface{}, get Getfunc, obsfunc Observer) {
	obs, _ := source.(Observable)
	if obs != nil {
		obs.AddObserver(obsfunc)
		return
	}
	// add to manual poll list instead
	if o.pollers == nil {
		o.pollers = make(map[interface{}]poll)
	}
	p := o.pollers[source]
	p.current = get()
	p.get = get

	p.AddObserver(obsfunc)
	o.pollers[source] = p
}

// Connects a source value with a target value. In other words, if
// the source value changes, the target will be set to match.
func (o *Observatory) Connect(source, target interface{}) {
	set, _ := target.(Setable)

	var (
		obsfunc func()
		get     = o.CreateGetfunc(source)
	)

	if set != nil {
		obsfunc = func() {
			set.Set(get())
		}
	} else if t := reflect.ValueOf(target); t.Kind() == reflect.Ptr {
		switch t := t.Interface().(type) {
		case (*float32):
			obsfunc = func() {
				*t = float32(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Float())
			}
		case (*float64):
			obsfunc = func() {
				*t = float64(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Float())
			}
		case (*int8):
			obsfunc = func() {
				*t = int8(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Int())
			}
		case (*int16):
			obsfunc = func() {
				*t = int16(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Int())
			}
		case (*int32):
			obsfunc = func() {
				*t = int32(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Int())
			}
		case (*int64):
			obsfunc = func() {
				*t = int64(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Int())
			}
		case (*int):
			obsfunc = func() {
				*t = int(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Int())
			}
		case (*uint8):
			obsfunc = func() {
				*t = uint8(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Uint())
			}
		case (*uint16):
			obsfunc = func() {
				*t = uint16(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Uint())
			}
		case (*uint32):
			obsfunc = func() {
				*t = uint32(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Uint())
			}
		case (*uint64):
			obsfunc = func() {
				*t = uint64(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Uint())
			}
		case (*uint):
			obsfunc = func() {
				*t = uint(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Uint())
			}
		case (*bool):
			obsfunc = func() {
				*t = bool(reflect.ValueOf(get()).Convert(reflect.TypeOf(*t)).Bool())
			}
		default:
			panic(fmt.Errorf("Can't deal with that non-setable target type: %s", reflect.TypeOf(target)))
		}
	} else {
		panic(fmt.Errorf("Can't deal with that non-setable target type: %s", reflect.TypeOf(target)))
	}
	o.Observe(source, get, funcObserver{obsfunc})
}

// Checks all poll sources for changes, and invokes
// the observers as needed.
func (o *Observatory) Update() {
	redo := true
	const maxLoops = 1000
	for i := 0; i < maxLoops && redo; i++ {
		redo = false
		for k, p := range o.pollers {
			nv := p.get()
			if nv == p.current {
				continue
			}
			p.current = nv
			o.pollers[k] = p
			p.NotifyObservers(nil)
			redo = true
		}
	}
}
