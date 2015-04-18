package main

import (
	"fmt"
	"math"
)

func main() {
	var r, g, b, f float64

	f = 0.1
	for i := 0; i < 360; i += 10 {
		r = math.Sin(f*float64(i)+0)*127 + 128
		g = math.Sin(f*float64(i)+2)*127 + 128
		b = math.Sin(f*float64(i)+4)*127 + 128
		fmt.Printf("%4v %4v %4v\n", r, g, b)
	}
}
