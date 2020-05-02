package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {

	case t.Hour() < 12 && t.Hour() >= 5:
		fmt.Println("Good morning!") //execute this if t.Hour() < 12 == true

	case t.Hour() < 17 && t.Hour() >= 12:
		fmt.Println("Good afternoon.")

	case t.Hour() >= 21 || t.Hour() < 5:
		fmt.Println("Good Night.")

	default:
		fmt.Println("Good evening.")

	}

}
