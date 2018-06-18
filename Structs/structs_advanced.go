package main

import "fmt"

// Foo is a struct here
type Foo struct {
	A int
	b string
}

// Bar is yet another struct that has F of Foo type
type Bar struct {
	C string
	F Foo
}

// Baz has a struct declaration, this will show us embedded delegation
type Baz struct {
	D string
	Foo
}

func main() {

	f := Foo{A:10, b:"Hey"}
	fmt.Println(f)

	g := Bar{C:"Hola", F:f}
	fmt.Println(g.F.A) // That's how you call values from that struct type var

	h := Baz{D:"Hi", Foo:f}
	fmt.Println(h.b) // That's how you call value from an embedded struct

	var f1 Foo
	f1 = h.Foo
	fmt.Println(f1)
}
