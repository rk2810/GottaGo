package main

import "fmt"

// Foo is a struct here, We dont have inheritance in Go, and hence structs are different than objects
type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	fmt.Println(g.b)
	// Ways of adding values to structs

	f.b = "Hey" // If you are not supplying value to one of the data type in struct, It automatically takes zero value
			   // for it.
	h := Foo{
		A : 20,
	}
	
	fmt.Println(h)

	fmt.Println(f)

	// Structs in Go are value type, example >

	var f2 Foo
	f2 = f
	f2.A = 100

	fmt.Println(f2.A)
	fmt.Println(f.A)

	// Although, we can use pointers as well

	var f3 *Foo
	f3 = &f
	f3.A = 69
	
	fmt.Println(f3.A) // Since f3 points to f, the value assignemnt reflects on f as well
	fmt.Println(f.A)

}
