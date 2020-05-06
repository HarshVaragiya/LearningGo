package main

import "fmt"

func fibonacci() func() int {

	a := 0
	b := 1
	var c int

	return func() int {

		oldA := a
		c = a + b
		a = b
		b = c
		return oldA

	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 25; i++ {
		fmt.Print(f(), ",")
	}

}
