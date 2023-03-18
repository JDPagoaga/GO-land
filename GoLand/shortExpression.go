package main

import "fmt"

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

}
