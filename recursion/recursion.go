package main

import "fmt"

func main() {
	factNumber := 5
	fact := fact(factNumber)
	fmt.Printf("factorial on %d is %d\n", factNumber, fact)

	s := [] int {3,2,8}
	sum := sum(s)
	fmt.Printf("sum on %v is %d\n", s, sum)

}

func fact(x int) int {
	if x == 1 {
		return x
	} else {
		return x * fact(x-1)
	}
}

func sum (s []int) int {
	if(len(s) == 0) {
		return 0
	}else {
		return s[len(s)-1] + sum(s[0:len(s)-1])
	}
}
