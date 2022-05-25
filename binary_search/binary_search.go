package main

import (
	"fmt"
	"errors"
	"log"
)

func main() {
	//TODO разобарться с make и new
	//TODO разобраться с cap и len,
	//TODO разобраться как создать слайс не указывая его длину и заполнить его (можно использовать аппенд, но тогда нужно возвращать массив, а нужно понять как менять исходный)
	//TODO разобраться как написать метод fillHaystack используя ресивер
	
	max := 10
	haystack := make([]int, max)
	fillHaystack(haystack, max)
	test,err := binary_search(7, haystack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", test)
}

func fillHaystack(haystack []int, max int) {
	for i := 0; i < max; i++ {
		haystack[i] = i+1
	}
}

func binary_search(needle int, haystack[]int) (int, error){
	
	for len(haystack) > 1 {
		fmt.Printf("len: %d\n", len(haystack))
		maxIndex := len(haystack) - 1 
		midIndex := len(haystack) / 2
		if haystack[midIndex] == needle {
			return midIndex, nil
		} else if haystack[midIndex] > needle{
			haystack = haystack[0:midIndex]
		} else {
			haystack = haystack[midIndex:maxIndex]
		}
	}

	return -1, errors.New("number not found")
}