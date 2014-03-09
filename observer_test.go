// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package util

import (
	"testing"
)

type aaa int

func (a *aaa) Set(v interface{}) {
	switch t := v.(type) {
	case (*int):
		*a = *(*aaa)(t)
	case (int):
		*a = aaa(t)
	case (*aaa):
		*a = *t
	case (aaa):
		*a = t
	}
}

func TestObseratory1(t *testing.T) {
	var (
		a = aaa(1)
		b = aaa(2)
		o Observatory
	)
	o.Connect(&a, &b)
	a.Set(9)
	o.Update()
	if a != 9 || b != 9 {
		t.Errorf("%d, %d", a, b)
	}
	b.Set(10)
	o.Update()
	if a != 9 || b != 10 {
		t.Errorf("%d, %d", a, b)
	}
	a.Set(11)
	//	o.Update()
	if a != 11 || b != 10 {
		t.Errorf("%d, %d", a, b)
	}
}

func TestObseratory2(t *testing.T) {
	var (
		a = 1
		b = 2
		o Observatory
	)
	o.Connect(&a, &b)
	a = 9
	o.Update()
	if a != 9 || b != 9 {
		t.Errorf("%d, %d", a, b)
	}
	b = 10
	o.Update()
	if a != 9 || b != 10 {
		t.Errorf("%d, %d", a, b)
	}
	a = 11
	//	o.Update()
	if a != 11 || b != 10 {
		t.Errorf("%d, %d", a, b)
	}
}

func TestObseratory3(t *testing.T) {
	var (
		a bool
		b bool
		o Observatory
	)
	o.Connect(&a, &b)
	a = true
	o.Update()
	if !a || a != b {
		t.Errorf("%v, %v", a, b)
	}
	b = false
	o.Update()
	if !a || b {
		t.Errorf("%v, %v", a, b)
	}
	b = true
	a = false
	//	o.Update()
	if a || !b {
		t.Errorf("%v, %v", a, b)
	}
}

func TestObseratory4(t *testing.T) {
	var (
		a int
		b float32
		o Observatory
	)
	o.Connect(&a, &b)
	a = 10
	o.Update()
	if a != int(b) {
		t.Errorf("%v, %v", a, b)
	}
	b = 3
	o.Update()
	if a == int(b) {
		t.Errorf("%v, %v", a, b)
	}
	a = 99
	//	o.Update()
	if a != 99 || b != 3 {
		t.Errorf("%v, %v", a, b)
	}
}

func TestObseratory5(t *testing.T) {
	var (
		a int
		b uint8
		o Observatory
	)
	o.Connect(&a, &b)
	a = 10
	o.Update()
	if a != int(b) {
		t.Errorf("%v, %v", a, b)
	}
	b = 3
	o.Update()
	if a == int(b) {
		t.Errorf("%v, %v", a, b)
	}
	a = 99
	//	o.Update()
	if a != 99 || b != 3 {
		t.Errorf("%v, %v", a, b)
	}
}
