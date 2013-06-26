package container

type (
	Lifo struct {
		Fifo
	}
)

func (li *Lifo) Less(i, j int) bool {
	return li.data[i].index > li.data[j].index
}
