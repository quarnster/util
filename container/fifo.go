package container

type (
	data struct {
		index int
		value interface{}
	}
	Fifo struct {
		counter int
		data    []data
	}
)

func (fi *Fifo) Len() int { return len(fi.data) }

func (fi *Fifo) Less(i, j int) bool {
	return fi.data[i].index < fi.data[j].index
}

func (fi *Fifo) Swap(i, j int) {
	fi.data[i], fi.data[j] = fi.data[j], fi.data[i]
}

func (fi *Fifo) Push(x interface{}) {
	fi.counter++
	fi.data = append(fi.data, data{fi.counter, x})
}

func (fi *Fifo) Pop() interface{} {
	n := len(fi.data)
	ret := fi.data[n-1].value
	fi.data = fi.data[:n-1]
	return ret
}
