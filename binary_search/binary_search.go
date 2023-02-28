package main

import (
	"fmt"
)

func main() {

	var max int
	fmt.Println("Enter max haystack size...")
	fmt.Scanf("%d\n", &max)

	var needle int
	fmt.Println("Enter needle...")
	fmt.Scanf("%d\n", &needle)

	var haystack []int
	FillHaystack(&haystack, max)
	fmt.Println(haystack)
	key := BinarySearch(needle, haystack)

	fmt.Println("binary search:")
	fmt.Printf("needle key: %d\n", key)

	key = StupidSearch(needle, haystack)
	fmt.Println("stupid search:")
	fmt.Printf("needle key: %d\n", key)
}

func FillHaystack(haystack *[]int, max int) {
	for i := 0; i < max; i++ {
		*haystack = append(*haystack, i+1)
	}
}

func BinarySearch(needle int, haystack []int) int {
	minIndex := 0
	maxIndex := len(haystack) - 1

	for minIndex <= maxIndex {
		midIndex := (minIndex + maxIndex) / 2
		if haystack[midIndex] == needle {
			return midIndex
		} else if haystack[midIndex] > needle {
			maxIndex = midIndex - 1
		} else {
			minIndex = midIndex + 1
		}
	}

	return -1
}

func StupidSearch(needle int, haystack []int) int {

	for key, val := range haystack {
		if val == needle {
			return key
		}
	}

	return -1
}
