package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Ask for the user's name using a print statement
	fmt.Print("Please Enter your name : ")

	// Take input from User as their name
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Display a message to the user using the input string
	fmt.Println("Hello", input[:len(input)-2], "nice to meet you..")

}
