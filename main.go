package main

import (
	"fmt"
	"math"
)

func rgb2ansi(r, g, b int) int {
	return 16 + 36*int(r*6/256) + 6*int(g*6/256) + int(b*6/256)
	//return 16 + 36*r*6/256 + 6*g*6/256 + b*6/256
	//return 16 + (36*r+6*g+b)*6/256
}

func rgb(r, g, b int) int {
	return 16 + 36*r + 6*g + b
}

func main() {
	var r, g, b, r5, g5, b5 int
	var f float64

	f = 0.1
	for i := 0; i < 50; i += 1 {
		r = int(math.Sin(f*float64(i)+0)*127 + 128)
		g = int(math.Sin(f*float64(i)+2)*127 + 128)
		b = int(math.Sin(f*float64(i)+4)*127 + 128)
		//fmt.Printf(`<font color="%02X%02X%02X">&#9608;</font>`, r, g, b)

		r5 = r * 6 / 256
		g5 = g * 6 / 256
		b5 = b * 6 / 256
		_ = r5 + g5 + b5
		//fmt.Println(r5, g5, b5, rgb(r5, g5, b5), rgb2ansi(r, g, b))

		//fmt.Println(rgb(r5, g5, b5))
		fmt.Println(rgb2ansi(r, g, b))
	}

}
