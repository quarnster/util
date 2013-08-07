package container

import (
	"math/rand"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	const count = 100
	for i := 0; i < 10; i++ {
		tree := Tree{Less: func(a, b interface{}) bool {
			return a.(int) < b.(int)
		}}
		list := rand.Perm(count)
		for _, j := range list {
			tree.Add(j)
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
