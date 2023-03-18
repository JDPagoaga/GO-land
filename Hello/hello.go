package main

import "fmt"

func main() {
	fmt.Println("Hello EveryBody, Im so Happy")

	Bye()

	fmt.Println("It's the end")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "bye")
		}
	}
}

func Bye() {
	fmt.Println("Well, see you later")
}
