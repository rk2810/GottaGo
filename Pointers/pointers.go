package main

import "fmt"

func makePoint(a *int){
	// a = new(int)  // Uncomment this and you will see magic of call by value concept in Go.
	*a = 40 // this is called value of a
}

func main() {
	a := 10
	b := &a // &a is called reference to a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	// var x *int
	x := new(int)
	fmt.Println(x)
	fmt.Println(*x)

	z := 20
	fmt.Println(z)
	makePoint(&z) // the function expected a pointer, so we pass reference to z in args
	fmt.Println(z)


}
