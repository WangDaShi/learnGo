package main

import (
	"fmt"
)

func main() {
	//loatr.TestChannel()
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

type sss struct {
	x, y int64
}

func (s sss) String() string {
	return fmt.Sprintln("x = %v,y = %v", s.x, s.y)
}
