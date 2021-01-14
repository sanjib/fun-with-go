package main

import "golang.org/x/tour/pic"

// Pic creates a picture
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, 0)

	for y := range make([]int, dy) {
		row := make([]uint8, 0)
		for x := range make([]int, dx) {
			// formula := uint8((x + y) / 2)
			// formula := uint8(x * y)
			formula := uint8(x ^ y)
			row = append(row, formula)
		}
		pic = append(pic, row)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
