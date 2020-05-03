package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var a point = point{1, 2}
	var b point = point{3, 4}

	p, q := &a, &b

	p.x = 5 // No need of (*p).x => directly work with pointer variable
	q.x = 5

	fmt.Println(a, b)

}
