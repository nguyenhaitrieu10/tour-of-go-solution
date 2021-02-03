// Exercise: Loops and Functions
package main

import (
	"fmt"
	"math"
)

const Epsilon = 0.00001

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		fmt.Println(z)
		tmp := z - (z*z-x)/(2*z)
		if math.Abs(tmp-z) < Epsilon {
			break
		}
		z = tmp
	}
	return z
}

func main() {
	fmt.Println(Sqrt(13689))
}
