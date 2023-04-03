package main

import (
	"container/heap"
	"fmt"
)

type (
	item struct {
		index    int
		priority int
	}
	PriorityQueue []item
)

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	count := len(old)
	x := old[count-1]
	*pq = old[0 : count-1]
	return x
}

func main() {
	h := &PriorityQueue{
		{
			index:    3,
			priority: 8,
		},
		{
			index:    2,
			priority: 5,
		},
		{
			index:    1,
			priority: 1,
		},
	}

	h.Push(item{
		index:    5,
		priority: 7,
	})
	heap.Init(h)
	for h.Len() > 0 {
		pop := heap.Pop(h)
		fmt.Println(pop.(item).index)
	}
}
