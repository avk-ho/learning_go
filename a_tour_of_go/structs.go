package main

import (
	"fmt"
	"math"
	"strings"
)

// STRUCT AND METHODS

type Vertex struct {
	X float64
	Y float64
}

// a method of Vertex, which is put after func, 
// and before the method name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function version of Abs
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// a method that uses a pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function version of Scale
// without a pointer, the operations are done on
// a copy of v
func ScaleFunc(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// the same function with pointer
func ScaleFuncPointer(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// functions with pointer MUST take a pointer
// methods with pointer can take either a value or a pointer
// function with a value MUST take a value
// methods with a value can take either a value or a pointer

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


// INTERFACE

// an interface contains a set of method signatures
// a value of the interface type can hold any value that implements
// those methods
// implementing an interface simply requires its methods to be implemented

type Abser interface {
	Abs() float64
}

// an interface with the method M
type I interface {
	M()
}

// a struct
type T struct {
	S string
}

// a method of the struct T
func (t *T) M() {
	// handling a nil value
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f MyFloat) M() {
	fmt.Println(f)
}

// demonstrate that an interface is similar to a tuple (value, type)
// and a interface value saves the underlying concrete type
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// TYPE SWITCH

func do(i interface{}) {
	// getting the type of i via the keyword type
	switch v := i.(type) {
	// checking for potential types:
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// STRINGER EXERCISE

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringerExercise() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


// MAIN

func main() {
	v := Vertex{1, 2}

	fmt.Println(v)

	// how to access struct fields
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// using the Abs method
	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())

	// using a method on a non-struct type
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// using the pointer method
	v.Scale(10)

	// function, without pointer
	v3 := Vertex{3, 4}
	ScaleFunc(v3, 10)
	fmt.Println(AbsFunc(v3))

	// function, with pointer
	ScaleFuncPointer(&v3, 10)
	fmt.Println(AbsFunc(v3))


	// implementing the interface Abser
	var a Abser
	it := MyFloat(-math.Sqrt2)
	a = it
	a = &v2

	fmt.Println(a.Abs())

	// implementing the interface I
	var i I = &T{"hello"}
	describe(i) // main T type
	i.M()

	i = MyFloat(math.Pi)
	describe(i) // main MyFloat type
	i.M()

	var t *T
	i = t
	describe(i) // value is nil, main T type
	i.M()

	// nil interface value
	var i2 I
	describe(i2) // value is nil, main nil type

	// will cause a runtime error, cannot call a method
	// on a nil interface value
	// i.M()


	// empty interface
	var i3 interface{}
	describeEmpty(i3)

	// an empty interface can hold values of any type and
	// is used by code that handles values of unknown type
	i3 = 42
	describeEmpty(i3)

	i3 = "hello"
	describeEmpty(i3)


	// type assertions
	// allows to access the underlying TYPE T of the interface value
	// i and assigns its VALUE to the variable t
	// t := i.(T)

	// if i does not hold a T, an error will be triggered
	// ok here will be true if T exists, else false
	// t, ok, := i.(T)
	fmt.Println("\nType assertions: ")
	s := i3.(string)
	fmt.Println(s)

	s, ok := i3.(string)
	fmt.Println(s, ok)

	f2, ok := i3.(float64)
	fmt.Println(f2, ok)

	// f2 = i3.(float64) // panic
	// fmt.Println(f2)


	fmt.Println("\nType switch: ")
	do(21)
	do("hello")
	do(true)

	fmt.Println("\nStringer Exercise: ")
	stringerExercise()

}
