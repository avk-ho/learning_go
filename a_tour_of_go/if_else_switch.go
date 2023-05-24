package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	// simple if
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if/else
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	// for loop
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		
		fmt.Println(z)
	}
	return z
}

func distanceToFriday() {
	// switch evaluation, top to bottom
	fmt.Println("When's Friday?")
	today := time.Now().Weekday()
	//fmt.Println(today)
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func greetings() {
	// switch with no condition
	// an alternative to if/elif/else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	fmt.Println("sqrt func: ")
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println("\npow func: ")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("\nSqrt func: ")
	fmt.Println(Sqrt(4))


	// switch
	fmt.Println("\nswitch example: ")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("\ndistanceToFriday func: ")
	distanceToFriday()

	fmt.Println("\ngreetings func: ")
	greetings()
}
