package main

import (
	"math"
)

func pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			t := math.Log(float64(y))
			pic[x][y] = uint8(float64(x) * t)
		}
	}
	return pic
}


