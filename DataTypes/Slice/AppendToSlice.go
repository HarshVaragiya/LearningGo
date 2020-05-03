package main

import "fmt"

func main() {
	var s = []int{0}
	printSlice(s)

	// append works on nil slices.
	s = append(s, 1)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 2)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 3, 4, 5)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
