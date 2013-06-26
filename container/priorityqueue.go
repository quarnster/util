package container

type PriorityQueue struct {
	Lifo
	Priority func(x interface{}) int
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.data = append(pq.data, data{pq.Priority(x), x})
}
