package main

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Init(list []int) {
	h.list = list
	for i := len(h.list) / 2; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *Heap) Add(item int) {
	h.list = append(h.list, item)
	if len(h.list) == 1 {
		return
	}

	i := len(h.list) - 1
	parent := (i - 1) / 2

	for h.list[parent] < h.list[i] {
		tmp := h.list[parent]
		h.list[parent] = h.list[i]
		h.list[i] = tmp

		i = parent
		parent = (i - 1) / 2
	}

}

func (h *Heap) Max() (int, bool) {
	if len(h.list) == 0 {
		return 0, false
	}

	return h.list[0], true
}

func (h *Heap) RemoveMax() (int, bool) {
	if len(h.list) == 0 {
		return 0, false
	}
	top := h.list[0]
	if len(h.list) > 1 {
		h.list = h.list[1:]
	} else {
		h.list = []int{}
	}
	h.heapify(0)
	return top, true
}

func (h *Heap) heapify(i int) {
	leftChild := (i * 2) + 1
	rightChild := (i * 2) + 2

	for {
		if leftChild > len(h.list)-1 {
			break // don't have leafs
		}

		var maxChildItem int
		var maxChild int
		if rightChild <= len(h.list)-1 {
			if h.list[leftChild] > h.list[rightChild] {
				maxChildItem = h.list[leftChild]
				maxChild = leftChild
			} else {
				maxChildItem = h.list[rightChild]
				maxChild = rightChild
			}
		} else {
			maxChildItem = h.list[leftChild]
			maxChild = leftChild
		}

		if h.list[i] < maxChildItem {
			tmp := h.list[maxChild]
			h.list[maxChild] = h.list[i]
			h.list[i] = tmp

			i = maxChild
			leftChild = (i * 2) + 1
			rightChild = (i * 2) + 2
		} else {
			break
		}
	}
}

func main() {
	heap := Heap{}
	sl := []int{-5, 2, 3, 4, 5, 12}
	heap.Init(sl)
	fmt.Println(heap.list)
	top, ok := heap.RemoveMax()
	fmt.Println(top, ok)
	fmt.Println(heap.list)
}
