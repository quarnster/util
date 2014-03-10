// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package util

import (
	"math/rand"
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

func TestObseratory6(t *testing.T) {
	var (
		values    = make([]int, 100)
		halfway   = len(values) / 2
		o         Observatory
		obsCount1 int
		obsCount2 int
	)
	conn := rand.Perm(len(values))
	for i := 0; i < len(values)-1; i++ {
		a, b := conn[i], conn[i+1]
		o.Connect(&values[a], &values[b])
	}
	values[conn[0]] = 1
	src := &values[conn[halfway-1]]
	o.Observe(src, o.CreateGetfunc(src), func() { obsCount1++ })
	src = &values[conn[halfway+1]]
	o.Observe(src, o.CreateGetfunc(src), func() { obsCount2++ })
	o.Update()
	for _, v := range values {
		if v != 1 {
			t.Errorf("expected 1 not %d", v)
		}
	}
	values[conn[halfway]] = 2
	o.Update()

	for i, v := range conn {
		v = values[v]
		exp := 1
		if i >= halfway {
			exp = 2
		}
		if v != exp {
			t.Errorf("expected %d not %d", exp, v)
		}
	}
	if obsCount1 != 1 || obsCount2 != 2 {
		t.Errorf("obsCount's are wrong: %d, %d", obsCount1, obsCount2)
	}
}
