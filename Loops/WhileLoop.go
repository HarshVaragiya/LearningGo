package main

import "fmt"

func main() {

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i+=1
	}

	fmt.Println("There isn't a while loop in Go!")
}
