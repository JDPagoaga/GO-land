package main

import "fmt"

type My_type int

var xx My_type

func main() {
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)
	xx = 42
	fmt.Println(xx)
}
