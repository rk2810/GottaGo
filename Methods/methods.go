package main

import "fmt"

// Foo is a struct with A int and B string
type Foo struct {
	A int
	B string
}

// Concept of methods, methods vs. functions
// Go does not have classes, yet we can use structs and pass them in functions

// String method has a reciever of type Foo named f
func (f Foo) String() string {
	return fmt.Sprintf("A : %d and B: %s", f.A, f.B)
}

// Double method has a reciever of type pointer to Foo named f
func (f *Foo) Double() {
	f.A = f.A * 2
}

func main() {
	//  Method is a function with a special reciever argument
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())
}
