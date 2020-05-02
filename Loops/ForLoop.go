package main

import "fmt"

func main() {
	var sum,i,sum_upto uint32= 0 ,0 ,10

	for i = 1 ; i <= sum_upto; i++ {

		sum += i
		fmt.Println("[DEBUG] i = ", i, " sum = ",sum)

	}

	fmt.Println("[INFO] Sum till ", sum_upto," = " , sum)
}
