package main

import "fmt"
/*
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
 */
func main() {

	var str string = "Hello World!"
	defer fmt.Println(str)
	str = "Main Ends Here..."
	fmt.Println(str)

}
