package main

import "fmt"

var g int = 42
var h string = "James Bond"
var j bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", g, h, j)
	fmt.Println(s)

}
