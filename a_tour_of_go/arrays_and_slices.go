package main

import (
	"fmt"
	// "golang.org/x/tour/pic"
)

func sliceModifications() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// modifying a slice of an array changes the array itself
	// and other slices of the same array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slicesCapLen() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Length and capacity are also available for arrays
	// s2 := [4]int{1, 2, 3, 4}
	// fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceMake() {
	// Create a slice of int, length 5
	a := make([]int, 5)
	printSliceMake("a", a)

	// Create a slice of int, length 0, capacity 5
	b := make([]int, 0, 5)
	printSliceMake("b", b)

	// Slice len 2, cap 5
	c := b[:2]
	printSliceMake("c", c)

	// Slice len 3, cap 3
	d := c[2:5]
	printSliceMake("d", d)
}

func sliceAppend() {
	var s []int
	printSlice(s)

	// append(slice, new elem), new elem must be of a valid type
	// if the original array is too small, a new bigger array
	// will be created, and the slices will point toward it

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := range result {
		result[y] = make([]uint8, dx)

		for x := range result[y] {
			// num := uint8(x*y)
			// num := uint8(x^y)
			num := uint8((x+y)/2)

			result[y][x] = num
		}

	}

	return result
}

func main() {
	// a is an array of string, of capacity 2
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	// accessing elements of a via index
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// primes is an array of ints, of capacity 6
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// s is a slice of primes
	// a[x:y] contains elems from index x (included) to y (non included)
	fmt.Println("\nSlice of primes: ")
	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("\nModification of slices and underlying array: ")
	sliceModifications()

	fmt.Println("\nChecking length and capacity of slices: ")
	slicesCapLen()

	fmt.Println("\nCreating slices with make(): ")
	sliceMake()

	fmt.Println("\nTesting slice append: ")
	sliceAppend()

	// https://go.dev/tour/moretypes/18
	// fmt.Println("\nPic exercise: ")
	// pic.Show(Pic)
}
