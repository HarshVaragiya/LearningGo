package main

import "fmt"

func probability(OutcomesA,OutcomesB float64)(Pa,Pb float64){
	var TotalOutcomes = OutcomesA + OutcomesB
	Pa = OutcomesA / TotalOutcomes
	Pb = OutcomesB / TotalOutcomes
	return
}

func main() {
	fmt.Println("Favourable Outcomes = 5 \nUnfavourable Outcomes = 5")
	Pa , Pb := probability(5,5)
	fmt.Printf("Pfavourable %g , Punfavourable %g",Pa,Pb)
}
