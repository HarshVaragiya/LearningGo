package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter your name : ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	fmt.Println("Input = \"", input, "\", Input Length = ", len(input))

	if input == "admin" {
		fmt.Println("Welcome Admin")
	} else if input == "harsh" {
		fmt.Println("Please login using you main account")
	} else {
		fmt.Println("Unauthorized!")
	}

}
