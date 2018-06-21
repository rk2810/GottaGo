package main

import "fmt"

// Yet another way to know the underlying concrete type of an interface

func doStuff(i interface{}) { // interface parameter
	switch i := i.(type){    // Playing with switch by keeping type a variable with help of interface
	case int:
		fmt.Println("Double is", i * 2)
	case string:
		fmt.Println("String type data,", len(i), "length long.")
	default:
		fmt.Println("Unknown type encountered.")
	}
}

func main(){

	doStuff(5)
	doStuff("Hello")
	doStuff(true)
}
