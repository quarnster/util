package container

import "fmt"

const (
	Less ComparisonResult = iota
	Equal
	Greater
)

type (
	ComparisonResult int
	Node             struct {
		Data     interface{}
		Children [2]*Node
	}
	Compare func(a, b interface{}) ComparisonResult
	Tree    struct {
		Compare Compare
		Root    Node
	}
)

func (n *Node) find(data interface{}, cmp Compare, child int, parent *Node) (rchild int, retparent, node *Node) {
	if n.Data == nil {
		return child, parent, n
	}

	switch c := cmp(data, n.Data); c {
	case Equal:
		return child, parent, n
	case Less:
		if n.Children[0] == nil {
			return 0, n, n.Children[0]
		} else {
			return n.Children[0].find(data, cmp, 0, n)
		}
	case Greater:
		if n.Children[1] == nil {
			return 1, n, n.Children[1]
		} else {
			return n.Children[1].find(data, cmp, 1, n)
		}
	default:
		panic(c)
	}
}

func (n *Node) Find(data interface{}, cmp Compare) (child int, parent, node *Node) {
	return n.find(data, cmp, -1, nil)
}

func (n *Node) Walk(ch chan interface{}) {
	if n.Children[0] != nil {
		n.Children[0].Walk(ch)
	}
	if n.Data != nil {
		ch <- n.Data
	}
	if n.Children[1] != nil {
		n.Children[1].Walk(ch)
	}
}

func (n *Node) delete(child int, parent *Node) {
	a, b := n.Children[0], n.Children[1]
	switch {
	case a == nil && b == nil:
		if parent != nil {
			parent.Children[child] = nil
		} else {
			n.Data = nil
		}
	case a == nil && b != nil:
		*n = *b
	case a != nil && b == nil:
		*n = *a
	default:
		n2 := b
		p2 := n
		for n2.Children[0] != nil {
			p2 = n2
			n2 = n2.Children[0]
		}
		n.Data = n2.Data
		if n2 == b {
			b.delete(1, n)
		} else {
			n2.delete(0, p2)
		}
	}
}

func (t *Tree) Find(data interface{}) (child int, parent, node *Node) {
	return t.Root.Find(data, t.Compare)
}

func (t *Tree) Add(data interface{}) error {
	child, p, n := t.Find(data)
	if n != nil {
		if n.Data == data {
			return fmt.Errorf("Data already exists in the tree")
		} else {
			n.Data = data
		}
	} else if p.Data != nil {
		p.Children[child] = &Node{Data: data}
	} else {
		panic("Both parent and child was null")
	}
	return nil
}

func (t *Tree) Delete(data interface{}) error {
	child, p, n := t.Find(data)
	if n == nil || (p == nil && n.Data == nil) {
		return fmt.Errorf("Unable to find that node")
	} else {
		n.delete(child, p)
		return nil
	}
}
