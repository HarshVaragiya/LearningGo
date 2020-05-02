package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// local variables

	const pi = 3.14159

	var (
		ToBe   bool       = false
		MaxInt uint32     = 1<<32 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	// printing types

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("Pi is :",pi)  // constant

}

/*

bool (defaults to false)

string (defaults to "")

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr  (all default to 0)

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

 */