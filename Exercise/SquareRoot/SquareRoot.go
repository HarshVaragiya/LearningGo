package main

import (
	"fmt"
	"math"
)

func display(z,diff float64){
	fmt.Println("[DEBUG] Z = ",z," Z*Z = ",z*z," Diff = ",diff)
}

func Sqrt(x float64)(z float64){

	z = 10089774
	diff := x - (z*z)
	display(z,diff)
	last_diff := diff
	i:=0
	const resolution float64 = 0.000001
	const MaxIters int = 0xffff

	for ;i< MaxIters;i++{

		z = z - ((z*z -x) / (2*z))

		last_diff = diff
		diff = x - (z*z)

		display(z,diff)

		if diff == 0 || last_diff == -diff || (math.Abs(last_diff) - math.Abs(diff) < resolution) {
			break
		}

	}

	fmt.Println("[DEBUG] Status:")
	fmt.Println("[DEBUG] Total Iterations Ran = ", i+1)
	fmt.Println("[INFO] Diff = ", diff)
	fmt.Println("[INFO] Sqrt = ", z)

	return
}

func main() {
	Sqrt(263647579680)
}
