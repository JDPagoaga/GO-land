package main

import "fmt"

//var keyword allow to declare variables out from main function
//be careful to declare variables in this way , it allow to inicialize but
//don't check the use
var m = "using var keyword"

// inicialize with default value
var b int
var c string
var flag bool

func main() {
	// with X : , declared the variable X
	// with =2 , assigned a value to variable X
	// with x := 42 , declare and assign to variable X at the same time
	x := 42
	fmt.Println(x)
	// with x = , assign a new value to variable x
	x = 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	y = y + x
	fmt.Println(y)
	z := "It's a string"
	fmt.Println(z)
	fmt.Println(m)
	fmt.Println("default value of int", b)
	fmt.Println("default value of string*", c, "*")
	fmt.Println("default value of boolean", flag)

}
