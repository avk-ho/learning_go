package main

import (
	"fmt"
)

func fizzbuzz() {
	// multiples of 3 fizz
	// multiples of 5 buzz
	for i := 1; i <= 100; i++ {
		toPrint := ""
		if i % 3 == 0 {
			toPrint += "fizz"
		}
		if i % 5 == 0 {
			toPrint += "buzz"
		}

		if toPrint == "" {
			fmt.Println(i)
		} else {
			fmt.Println(toPrint)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
