package main

import "fmt"

var p = true

type flagY bool
type flagN bool
type Cont int

var y flagY
var z flagY
var n flagN
var i Cont

func main() {
	fmt.Println("using several types")
	fmt.Println("print p ", p)
	fmt.Printf("%T\n", p)
	fmt.Println("print y ", y)
	fmt.Printf("%T\n", y)
	fmt.Println("print z ", z)
	fmt.Printf("%T\n", z)
	fmt.Println("print n ", n)
	fmt.Printf("%T\n", n)
	fmt.Println("print i ", i)
	fmt.Printf("%T\n", i)

	y = true
	//this action only can be, b'cause they are the same type
	z = y
	fmt.Println("print z ", z)
	// using conversions
	fmt.Println("using conversions ")
	p = bool(n)
	fmt.Println("new value of p", p)
	p = bool(y)
	fmt.Println("next value of p", p)
	z = flagY(n)
	fmt.Println("new  value of z", z)

}
