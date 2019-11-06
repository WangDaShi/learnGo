package loatr

import "fmt"

type emptyIf interface {
}

// TestAssert something
func TestAssert() {
	var i interface{} = "123"
	t := i.(string)
	fmt.Println(t)
	o := i.(int8)
	fmt.Println(o)
}
