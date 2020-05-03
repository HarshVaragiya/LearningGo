package main

import "fmt"

func main() {

	a, b := 10, 12
	fmt.Println("Values of a,b = ", a, b)
	x, y := &a, &b
	fmt.Println("Pointers to a,b = ", x, y)

	*x = *x + 2 // *x is same as a
	*y = *y / 2 // *y is same as b
	fmt.Println("Values of a,b = ", a, b)
	fmt.Println("Pointers to a,b = ", x, y)

}
