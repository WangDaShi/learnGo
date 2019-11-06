package loatr

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func TestInterface() {
	var abser Abser
	var ttt *vertex
	abser = ttt
	fmt.Println(abser.Abs())
	abser = Myfloat(12)
	fmt.Println(abser.Abs())
	v := vertex{3, 4}
	//abser = v
	//fmt.Println(abser.Abs())
	abser = &v
	fmt.Println(abser.Abs())
}

type Myfloat float64

func (m Myfloat) Abs() float64 {
	if m < 0 {
		return float64(m)
	}
	return float64(m)
}

type vertex struct {
	x, y float64
}

func (v *vertex) Abs() float64 {
	if v == nil {
		fmt.Println("the interface is nil")
		return 0
	}
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
