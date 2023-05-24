package main

import (
	"fmt"
	"math"
)

// it's possible to pass other functions as arguments
// or return values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// return a func
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	next_fib := 1
	
	return func() int {
		current_fib := fib

		fib = next_fib
		next_fib = current_fib + next_fib

		return current_fib
	}
}

func printFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("\nClosures: ")
	closure()

	fmt.Println("\nExercise Fibonacci: ")
	printFibonacci()
}
