package main

import "testing"

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := BinarySearch(3, haystack)
	if key != 2 {
		t.Errorf("BinarySearch(3, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]) = %v; want 2", key)
	}
}

func TestBinarySearch2(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := BinarySearch(11, haystack)
	if key != -1 {
		t.Errorf("BinarySearch(11, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]) = %v; want -1", key)
	}
}

func TestStupidSearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := StupidSearch(3, haystack)
	if key != 2 {
		t.Errorf("stupidSearch(3, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]) = %v; want 2", key)
	}
}

func TestStupidSearch2(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := StupidSearch(11, haystack)
	if key != -1 {
		t.Errorf("StupidSearch(11, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]) = %v; want -1", key)
	}
}
