package loatr

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

// TestGo l;k;k;
func TestGo() {
	go say("world")
	//say("hello")
	fmt.Println("hello")
	time.Sleep(1000 * 10 * time.Millisecond)
}
