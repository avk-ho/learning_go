package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)


type Vertex struct {
	Lat, Long float64
}

// maps are key-value pairs, m is a string-Vertex struct pair
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// no need to precise the type of the value here
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func modifyMap() {
	ma := make(map[string]int)

	// Inserting a pair
	ma["Answer"] = 42
	fmt.Println("The value:", ma["Answer"])

	// Accessing a value with the map and key
	elem := ma["Answer"]
	fmt.Println("elem value:", elem)

	// Modifying a value
	ma["Answer"] = 48
	fmt.Println("The value:", ma["Answer"])
	fmt.Println("elem value:", elem) // elem isn't changed

	// Deleting a value
	delete(ma, "Answer")
	fmt.Println("The value:", ma["Answer"])

	// checking that the value exists, returns the value and a bool
	v, ok := ma["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	// https://pkg.go.dev/strings#Fields
	s_slice := strings.Fields(s)
	var m = map[string]int{}

	for _, word := range(s_slice) {
		count, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word] = count + 1
		}
		fmt.Println("key: ", word, "value: ", m[word])
	}


	return m
}

func main() {
	fmt.Println("Printing a map: ")
	fmt.Println(m)

	fmt.Println("\nModifying maps: ")
	modifyMap()

	fmt.Println("\nExercise word count: ")
	wc.Test(WordCount)
}