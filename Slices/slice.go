package main

import (
	"fmt"
)

func main() {
	// A slice is a growable sequence of values of a single specified type.
	// Kinda growable array.

	s := make([]string, 0) // make function creates an instance for data types in go
	// we just made an instance for a slice 

	fmt.Println("length of s", len(s))
	s = append(s, "hello")
	s = append(s, "goodbye")
	fmt.Println(s)
	s[0] = "Hi"
	fmt.Println(s)
	fmt.Println("lenght of s", len(s))
	
	// using range to traverse slice
	for i := range s {
		fmt.Println("At index", i , "value is", s[i])
	}

}