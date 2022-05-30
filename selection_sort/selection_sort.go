package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {

	var max int 
	fmt.Println("Enter max slice size...")
	fmt.Scanf("%d\n", &max)

	randSlice := getRandSlice(max)

	fmt.Print("[")
	for _, r := range randSlice {
		// fmt.Printf("number: %d\n", r)
		fmt.Print(r, ", ")
	
	}
	fmt.Print("]\n")

	smallest := smallest(randSlice)
	fmt.Printf("Smallest value: %d\n", smallest)

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

func smallestIndex(s []int) int{
	smallest := s[0]
	smallestIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] < smallest {
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selection_sort(s []int) {
	var sortSlice []int
	for i := 0; i < len(s); i++ {
		smallestIndex := smallest(s)
		sortSlice = append(s, s[smallestIndex])
		
	}

}

func remove(s []int, i int) {
	
}