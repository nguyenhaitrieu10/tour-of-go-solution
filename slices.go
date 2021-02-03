// Exercise: Slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
		for j := range result[i] {
			// result[i][j] = uint8((j + i) / 2)
			// result[i][j] = uint8((j * i))
			result[i][j] = uint8((j ^ i))
		}

	}
	return result
}

func main() {
	pic.Show(Pic)
}
