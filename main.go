package main

import (
	"fmt"
	"math"
)

var f float64

func rgb2ansi(r, g, b int) int {
	return 16 + 36*r + 6*g + b
}

func lolrgb(i int) (int, int, int) {
	return int(math.Sin(f*float64(i)+0)*2.5 + 2.5),
		int(math.Sin(f*float64(i)+2)*2.5 + 2.5),
		int(math.Sin(f*float64(i)+4)*2.5 + 2.5)
}

func main() {
	var r, g, b int

	f = 0.1
	for i := 0; i < 50; i += 1 {
		r, g, b = lolrgb(i)
		//fmt.Printf(`<font color="%02X%02X%02X">&#9608;</font>`, r, g, b)

		//fmt.Println(rgb2ansi(r, g, b))
		fmt.Printf("\033[38;5;%dmo", rgb2ansi(r, g, b))
	}

}
