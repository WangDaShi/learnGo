package main

import (
	"fmt"

	"../loatr"
)

func main() {
	loatr.TestGo()
}

type sss struct {
	x, y int64
}

func (s sss) String() string {
	return fmt.Sprintln("x = %v,y = %v", s.x, s.y)
}
