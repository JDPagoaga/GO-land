package main

import "fmt"

type My_type2 int

var xx2 My_type2
var yy int

func main() {
	fmt.Println(xx2)
	fmt.Printf("%T\n", xx2)
	xx2 = 42
	fmt.Println(xx2)
	yy = int(xx2)
	fmt.Println(yy)
	fmt.Printf("%T\n", yy)

}
