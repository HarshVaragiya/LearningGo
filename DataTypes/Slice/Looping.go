package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func main() {

	for index, value := range pow {

		fmt.Printf("2^%d = %d\n", index, value)

	}

}
