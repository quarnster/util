// Copyright 2014 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package container_test

import (
	"github.com/quarnster/util/container"
	"testing"
)

var data = []interface{}{
	1, 2, 3, 4, "hello", "world",
}

func TestBasicArray_PushBack(t *testing.T) {
	a := &container.BasicArray{}
	// Insert data at end aka "push_back"
	for i, v := range data {
		if err := a.Insert(i, v); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		}
	}
	if l := a.Len(); l != len(data) {
		t.Errorf("Expected %d but got %d", len(data), l)
	}
	for i, v := range data {
		if v2 := a.Get(i); v != v2 {
			t.Errorf("%d: Expected %v, but got %v", i, v, v2)
		}
	}
}

func TestBasicArray_PopBack(t *testing.T) {
	a := &container.BasicArray{}
	for i, v := range data {
		if err := a.Insert(i, v); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		}
	}

	// Remove data at end aka "pop_back"
	for i := range data {
		i = len(data) - 1 - i
		v := data[i]
		if v2, err := a.Remove(i); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		} else if v != v2 {
			t.Errorf("%d: Expected %v, but got %v", i, v, v2)
		}
		if l := a.Len(); l != i {
			t.Errorf("%d: Expected Len %d, but got %d", i, i, l)
		}
	}
}

func TestBasicArray_PopFront(t *testing.T) {
	a := &container.BasicArray{}
	for i, v := range data {
		if err := a.Insert(i, v); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		}
	}

	// Remove data at front aka "pop_front"
	for i, v := range data {
		i2 := len(data) - 1 - i
		if v2, err := a.Remove(0); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		} else if v != v2 {
			t.Errorf("%d: Expected %v, but got %v", i, v, v2)
		}
		if l := a.Len(); l != i2 {
			t.Errorf("%d: Expected Len %d, but got %d", i, i2, l)
		}

	}
}

func TestBasicArray_PushFront(t *testing.T) {
	a := &container.BasicArray{}
	// Insert data at front aka "push_front" (i.e data should be reverse)
	for i, v := range data {
		if err := a.Insert(0, v); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		}
	}
	if l := a.Len(); l != len(data) {
		t.Errorf("Expected %d but got %d", len(data), l)
	}
	for i, v := range data {
		i2 := len(data) - 1 - i
		if v2 := a.Get(i2); v != v2 {
			t.Errorf("%d: Expected %v, but got %v", i, v, v2)
		}
	}

	for i, v := range data {
		i2 := len(data) - 1 - i
		if v2, err := a.Remove(i2); err != nil {
			t.Errorf("%d: Didn't expect an error but got one: %s", i, err)
		} else if v != v2 {
			t.Errorf("%d: Expected %v, but got %v", i, v, v2)
		}
		if l := a.Len(); l != i2 {
			t.Errorf("%d: Expected Len %d, but got %d", i, i2, l)
		}

	}
}
func TestIntArray(t *testing.T) {
	a := &container.IntArray{}
	exps := []bool{true, true, true, true, false, false}
	for i, d := range data {
		if err := a.Insert(i, d); (err == nil) != exps[i] {
			t.Errorf("%d, Error expectation mismatch. Expected an error? %v, got %s", i, exps[i], err)
		}
	}
	if l := a.Len(); l != 4 {
		t.Errorf("Expected %d but got %d", 4, l)
	}
}

func TestFilteredArray(t *testing.T) {
	var inner container.Array = &container.BasicArray{}
	check := func(data interface{}) bool {
		_, ok := data.(string)
		return ok
	}
	a, err := container.NewFilteredArray(inner, check)
	if err == nil {
		t.Errorf("Expected an error but didn't get one")
	}
	inner = &container.ObservableArray{Array: inner}
	a, err = container.NewFilteredArray(inner, check)
	if err != nil {
		t.Errorf("Didn't expect an error but got %s", err)
	}

	for _, d := range data {
		if err := inner.Insert(inner.Len(), d); err != nil {
			t.Errorf("Didn't expect an error but got %s", err)
		}
	}
	verify := func() {
		if l := a.Len(); l != 2 {
			t.Errorf("Expected %d but got %d", 2, l)
		}
		if v := a.Get(0); v != "hello" {
			t.Errorf("Expected %s but got %s", "hello", v)
		}
		if v := a.Get(1); v != "world" {
			t.Errorf("Expected %s but got %s", "world", v)
		}
	}
	verify()
	inner.Remove(0)
	verify()
	inner.Remove(0)
	verify()
	inner.Remove(inner.Len() - 2)
	if l := a.Len(); l != 1 {
		t.Errorf("Expected %d but got %d", 1, l)
	}
	if v := a.Get(0); v != "world" {
		t.Errorf("Expected %s but got %v", "world", v)
	}

}
