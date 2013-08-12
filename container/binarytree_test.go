package container

import (
	"math/rand"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	const count = 100
	for i := 0; i < 10; i++ {
		tree := Tree{Compare: func(a, b interface{}) ComparisonResult {
			aa := a.(int)
			bb := b.(int)
			switch {
			case aa < bb:
				return Less
			case aa > bb:
				return Greater
			default:
				return Equal
			}
		}}
		list := rand.Perm(count)
		for _, j := range list {
			if e := tree.Add(j); e != nil {
				t.Error(e)
			}
		}

		ch := make(chan interface{})
		go func() {
			tree.Root.Walk(ch)
			close(ch)
		}()
		for j := 0; j < count; j++ {
			k := (<-ch).(int)
			if k != j {
				t.Errorf("%d != %d", k, j)
			}
		}
	}
}

func TestBinaryTreeFind(t *testing.T) {
	const count = 100
	for i := 0; i < 10; i++ {
		tree := Tree{Compare: func(a, b interface{}) ComparisonResult {
			aa := a.(int)
			bb := b.(int)
			switch {
			case aa < bb:
				return Less
			case aa > bb:
				return Greater
			default:
				return Equal
			}
		}}
		list := rand.Perm(count)
		for _, j := range list {
			if e := tree.Add(j); e != nil {
				t.Error(e)
			}
		}
		list = rand.Perm(count)
		for _, j := range list {
			if _, _, n := tree.Find(j); n == nil {
				t.Errorf("Should have found %d, but didn't", j)
			} else if v, ok := n.Data.(int); !ok {
				t.Errorf("Unable to cast data to int... %+v", n.Data)
			} else if v != j {
				t.Errorf("Expected to find %d, but got %d", j, v)
			}
		}
	}
}

func TestBinaryTreeDelete(t *testing.T) {
	const count = 100
	const sub = 20
	for i := 0; i < 10; i++ {
		tree := Tree{Compare: func(a, b interface{}) ComparisonResult {
			aa := a.(int)
			bb := b.(int)
			switch {
			case aa < bb:
				return Less
			case aa > bb:
				return Greater
			default:
				return Equal
			}
		}}
		list := rand.Perm(count)
		for _, j := range list {
			if e := tree.Add(j); e != nil {
				t.Error(e)
			}
		}
		for j := 0; j < sub; j++ {
			if e := tree.Delete(j); e != nil {
				t.Error(e)
			}
		}

		ch := make(chan interface{})
		go func() {
			tree.Root.Walk(ch)
			close(ch)
		}()
		for j := 0; j < count-sub; j++ {
			k := (<-ch).(int)
			if k != (j + sub) {
				t.Errorf("%d != %d", k, j+sub)
			}
		}
	}
}

func TestBinaryTreeDelete2(t *testing.T) {
	const count = 100
	const sub = 20
	for i := 0; i < 10; i++ {
		tree := Tree{Compare: func(a, b interface{}) ComparisonResult {
			aa := a.(int)
			bb := b.(int)
			switch {
			case aa < bb:
				return Less
			case aa > bb:
				return Greater
			default:
				return Equal
			}
		}}
		list := rand.Perm(count)
		for _, j := range list {
			if e := tree.Add(j); e != nil {
				t.Error(e)
			}
		}
		list = rand.Perm(sub)
		for _, j := range list {
			if e := tree.Delete(j); e != nil {
				t.Error(e)
			}
		}

		ch := make(chan interface{})
		go func() {
			tree.Root.Walk(ch)
			close(ch)
		}()
		for j := 0; j < count-sub; j++ {
			k := (<-ch).(int)
			if k != (j + sub) {
				t.Errorf("%d != %d", k, j+sub)
			}
		}
	}
}

func TestBinaryTreeAddDelete(t *testing.T) {
	const count = 100
	const sub = 20
	for i := 0; i < 10; i++ {
		tree := Tree{Compare: func(a, b interface{}) ComparisonResult {
			aa := a.(int)
			bb := b.(int)
			switch {
			case aa < bb:
				return Less
			case aa > bb:
				return Greater
			default:
				return Equal
			}
		}}
		list := rand.Perm(count)
		for _, j := range list {
			if e := tree.Add(j); e != nil {
				t.Error(e)
			}
		}
		for k := 0; k < 10; k++ {
			list = rand.Perm(sub)
			a := rand.Intn(count - sub)
			for _, j := range list {
				if e := tree.Delete(a + j); e != nil {
					t.Error(e)
				}
			}
			list = rand.Perm(sub)
			for _, j := range list {
				if e := tree.Add(a + j); e != nil {
					t.Error(e)
				}
			}
		}

		ch := make(chan interface{})
		go func() {
			tree.Root.Walk(ch)
			close(ch)
		}()
		for j := 0; j < count; j++ {
			k := (<-ch).(int)
			if k != j {
				t.Errorf("%d != %d", k, j)
			}
		}
	}
}
