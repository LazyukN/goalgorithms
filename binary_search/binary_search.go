package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	var max int
	fmt.Println("Enter max haystack size...")
	fmt.Scanf("%d\n", &max)

	var needle int
	fmt.Println("Enter needle...")
	fmt.Scanf("%d\n", &needle)

	// haystack := make([]int, max)
	var haystack []int
	fillHaystack(&haystack, max)
	fmt.Println(haystack)
	key, err := binary_search(needle, haystack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("needle key: %d\n", key)
}

func fillHaystack(haystack *[]int, max int) {
	for i := 0; i < max; i++ {
		*haystack = append(*haystack, i+1)
	}
}

func binary_search(needle int, haystack []int) (int, error) {
	minIndex := 0
	maxIndex := len(haystack) - 1

	for minIndex <= maxIndex {
		midIndex := (minIndex + maxIndex) / 2
		if haystack[midIndex] == needle {
			return midIndex, nil
		} else if haystack[midIndex] > needle {
			maxIndex = midIndex - 1
		} else {
			minIndex = midIndex + 1
		}
	}

	return -1, errors.New("number not found")
}
