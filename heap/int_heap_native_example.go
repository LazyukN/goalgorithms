package main

import (
	"container/heap"
	"fmt"
)

// IntHeap это минимальная куча целых чисел.
type IntHeap []int

// Len, Less, Swap для реализации интерфейса sort.Interface
func (h IntHeap) Len() int { return len(h) }

// для реализации maxheap нужно инвертнуть метод Less
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push и Pop используют приемники указателей,
	// потому что они изменяют длину среза,
	// не только его содержимое.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Этот пример вставляет несколько int в IntHeap,
// проверяет минимум,
// и удаляет их в порядке приоритета.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
