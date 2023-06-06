package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	// ?
	currentCost := 0.0
	for maxMes := 0; ; maxMes++ {
		currentCost += 1 + (float64(maxMes) * 0.01)
		if currentCost >= thresh {
			return maxMes
		}
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}
