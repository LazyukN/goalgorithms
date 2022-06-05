package main 

import "fmt"


func main() {
	var number int
	fmt.Println("Enter number for factorial...")
	fmt.Scanf("%d\n", &number)

	fact := fact(number)
	fmt.Println(fact)
}


func fact(x int) int{
	if x == 1 {
		return x
	} else {
		return x * fact(x-1)
	}
}