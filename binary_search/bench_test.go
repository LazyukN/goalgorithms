package main

import "testing"

func BenchmarkBsearch(b *testing.B) {
	var haystack []int
	FillHaystack(&haystack, 1000000)
	for i := 0; i < b.N; i++ {
		BinarySearch(500, haystack)
	}
}

func BenchmarkSsearch(b *testing.B) {
	var haystack []int
	FillHaystack(&haystack, 1000000)
	for i := 0; i < b.N; i++ {
		StupidSearch(500, haystack)
	}
}
