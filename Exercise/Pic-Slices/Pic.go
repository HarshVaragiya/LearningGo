package main

import (
	"golang.org/x/tour/pic"
	"math"
)

const algo string = "mul2"

func transform(x, y uint8) (val uint8) {
	switch algo {
	case "avg":
		val = (x + y) / 2
	case "gradient1":
		val = y
	case "gradient2":
		val = x
	case "mul":
		val = x * y
	case "exp":
		val = uint8(math.Pow(float64(x), float64(y)))
	case "mul2":
		val = (x << 1) * (y >> 1)

	}
	return
}

func Pic(dx, dy int) [][]uint8 {
	var Picture = make([][]uint8, dy, dy)

	for i := dy; i > 0; i-- {
		var xAxis = make([]uint8, dx, dx)
		for j := dx; j > 0; j-- {
			xAxis[j-1] = transform(uint8(i), uint8(j))
		}
		Picture[i-1] = xAxis
	}

	return Picture
}

func main() {
	pic.Show(Pic)
}
