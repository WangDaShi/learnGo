package main

import "fmt"

type Sss struct {
	x, y float32
}

func main() {

	// var m map[string]Sss
	// m = make(map[string]Sss)
	// m["value"] = Sss{
	// 	123, 321,
	// }
	// fmt.Println(m["value"])

	m := make(map[string]string)
	m["Value"] = "key"
	fmt.Println(m["Value"])
}
