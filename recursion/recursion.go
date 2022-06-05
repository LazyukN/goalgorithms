package main

import "fmt"

func main() {
	factNumber := 5
	fact := fact(factNumber)
	fmt.Printf("factorial for %d is %d\n", factNumber, fact)

	s := []int{11}

	sum := sum(s)
	fmt.Printf("sum for %v is %d\n", s, sum)

	count := count(s)
	fmt.Printf("count for %v is %d\n", s, count)

	max := max(s)
	fmt.Printf("max for %v is %d\n", s, max)

}

func fact(x int) int {
	if x == 1 {
		return x
	}
	return x * fact(x-1)

}

func sum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1] + sum(s[0:len(s)-1])
}

func count(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return 1 + count(s[0:len(s)-1])
}

func max(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	if len(s) == 2 {
		if s[0] > s[1] {
			return s[0]
		} else {
			return s[1]
		}
	}
	max := max(s[0 : len(s)-1])
	if s[len(s)-1] > max {
		return s[len(s)-1]
	}

	return max
}
