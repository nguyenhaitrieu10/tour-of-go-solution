// Exercise: Errors
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

const Epsilon = 0.00001

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x / 2
	for {
		tmp := z - (z*z-x)/(2*z)
		if math.Abs(tmp-z) < Epsilon {
			break
		}
		z = tmp
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
