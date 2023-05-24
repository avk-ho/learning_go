package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range iterate over a slice or a map
	// for a slice, it returns 2 values per iteration:
	// idx and value (i and v here)
	// possibility of using _ if we want to ignore one or the other
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
