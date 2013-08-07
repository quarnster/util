package container

type (
	Node struct {
		Data     interface{}
		Children []Node
	}
	Less func(a, b interface{}) bool
	Tree struct {
		Less Less
		Root Node
	}
)

func (n *Node) Find(data interface{}, less Less) *Node {
	if n.Data == nil {
		return n
	} else {
		if n.Children == nil {
			n.Children = make([]Node, 2)
		}
		if less(data, n.Data) {
			return n.Children[0].Find(data, less)
		} else {
			return n.Children[1].Find(data, less)
		}
	}
}

func (n *Node) Walk(ch chan interface{}) {
	if n.Children != nil {
		n.Children[0].Walk(ch)
	}
	if n.Data != nil {
		ch <- n.Data
	}
	if n.Children != nil {
		n.Children[1].Walk(ch)
	}
}

func (t *Tree) Add(data interface{}) {
	n := t.Root.Find(data, t.Less)
	n.Data = data
}
