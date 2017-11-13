package main

import "golang.org/x/tour/pic"

//Function pic creates a  slice (with len dy) of slices of uints (with len dx) and then returns it
func Pic(dx, dy int) [][]uint8 {
	dyS := make([][]uint8, dy)
	for i, _ := range dyS {
		dyS[i] = make([]uint8, dx)
		for j, _ := range dyS[i] {
			dyS[i][j] = uint8((i + j) / 2)
		}
	}
	return dyS
}

func main() {
	// pic.Show IN: f func(int, int) [][]uint8
	pic.Show(Pic)
}
