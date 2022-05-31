package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var max int
	fmt.Println("Enter max slice size...")
	fmt.Scanf("%d\n", &max)

	randSlice := getRandSlice(max)
	fmt.Printf("rand: %v\n", randSlice)

	sortSlice := selectionSort(randSlice)
	fmt.Printf("sort: %v\n", sortSlice)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandSlice(max int) []int {
	randSlice := make([]int, max)
	for i := 0; i < len(randSlice); i++ {
		randSlice[i] = rand.Intn(len(randSlice))
	}

	return randSlice
}

func smallestIndex(s []int) int {
	smallest := s[0]
	smallestIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(s []int) []int {
	var sortSlice []int
	len := len(s)
	for i := 0; i < len; i++ {
		smallestIndex := smallestIndex(s)
		sortSlice = append(sortSlice, s[smallestIndex])
		s = remove(s, smallestIndex)
	}

	return sortSlice
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[0 : len(s)-1]
}
