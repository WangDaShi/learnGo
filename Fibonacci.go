package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	return func() (i int) {
		i = a
		b = a + b
		a = b - a
		return
	}
}

func print() {
	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
