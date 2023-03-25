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
	fmt.Println("print y ", y)
	fmt.Println("print z ", z)
	fmt.Println("print n ", n)
	fmt.Println("print i ", i)

	y = true
	//this action onli can be, b'cause they are the same type
	z = y
	fmt.Println("print z ", z)

}
