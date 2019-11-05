package loatr

import (
	"fmt"
	"math"
)

type sss struct {
	x, y float64
}

func (s sss) abs() float64 {
	return math.Sqrt(s.x*s.x + s.y*s.y)
}

func abs2(s sss) float64 {
	return math.Sqrt(s.x*s.x + s.y*s.y)
}

func (s *sss) scale(x float64) {
	s.x = s.x * x
	s.y = s.y * x
}

func scale2(s *sss, x float64) {
	s.x = s.x * x
	s.y = s.y * x
}

// Print they say this method mush have an comment
func Print() {
	s := sss{3, 4}
	var a *sss = &s
	a.scale(10)
	fmt.Println(s.abs())
	s.scale(10)
	(&s).scale(10)
	scale2(&s, 10)
	fmt.Println(abs2(s))
}
