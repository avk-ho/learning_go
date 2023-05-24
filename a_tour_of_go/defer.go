package main

import (
	"fmt"
)

func helloWorld() {
	// deferred statement are called when the function returns
	defer fmt.Println("world")

	fmt.Println("hello")
}

func chainedDefers() {
	// multiple defers are in a stack (last in, first out)
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	helloWorld()

	chainedDefers()
}