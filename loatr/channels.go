package loatr

import "fmt"

func sum(arr []int, c chan int) {
	var sum int = 0
	for _, i := range arr {
		sum += i
	}
	c <- sum
}

func TestChannel() {
	arr := []int{1, 2, 3, 4, 5, 6, 6, 8}
	c := make(chan int)
	d := make(chan int)
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], d)
	x, y := <-d, <-c
	fmt.Printf("x= %v,y = %v", x, y)
}
