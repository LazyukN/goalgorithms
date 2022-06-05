package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	s := []int{1, 3, 5, 2, 4, 0}
	fmt.Println("source: ", s)

	sortSlice := qsort(s)
	fmt.Println("sort:", sortSlice)
}

func qsort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	pivotIndex := rand.Intn(len(s))
	pivot := s[pivotIndex]

	var less []int
	var more []int
	for index, val := range s {
		if index == pivotIndex {
			continue
		}
		if val <= pivot {
			less = append(less, val)
		} else if val > pivot {
			more = append(more, val)
		}
	}

	sortLess := qsort(less)
	sortMore := qsort(more)
	less = append(sortLess, pivot)
	less = append(less, sortMore...)

	return less

}
