package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var a point = point{1, 2} // specify both x and y values
	var b point = point{x: 2} // specify only x value, y defaults to 0 as not specified
	var c point = point{}     // both default to 0

	fmt.Println(a.x, a.y)
	fmt.Println(b.x, b.y)
	fmt.Println(c.x, c.y)
}
