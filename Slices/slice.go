package main

import (
	"fmt"
)

func main() {
	// A slice is a growable sequence of values of a single specified type.
	// Kinda growable array.

	s := make([]string, 0) // make function creates an instance for data types in go
	// we just made an instance for a slice, here 0 is the current length of slice

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

	// What if we give a length while initializing a slice ?
	s1 := make([]string, 2) // length is 2
	fmt.Println("length of s1 is", len(s1))

	for i := range s1 {
		fmt.Println("At index", i , "value is", s1[i]) // so it basically gave the first two index what we call
		// zero value for string and thus it prints nothing
		
	}
	
	s1 = append(s1, "hello")
	fmt.Println(s1)
	fmt.Println(len(s1))

	// Another way of making a slice
	s2 := []string{"a", "b", "c"}

	for k,v := range s2{
		fmt.Println(k,v)
	}

	// Creating Slice from a slice
	s3 := s2[0:3] // if you are confused why 0:3, uncomment the for loop below and OBSERVE
	fmt.Println(s3)
	s3 = append(s3, "d")
	fmt.Println("Length of s3 is", len(s3), "and values are ", s3)
	s2[0] = "x"
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("Length of s3 is", len(s3), "and values are ", s3)

	// for x := 0; x<3; x++ {
	// 	fmt.Println(x)
	// }

	// Yet another way of  devlaring a slice

	var s4 []string
	s4 = s3
	s4[0] = "Camel"
	fmt.Println(s3)
	s4 = append(s4, "dog")
	fmt.Println("S3 is", s3)
	fmt.Println("S4 is", s4)

	// Strings and slices

	new := "🌍"
	bytes := []byte(new)
	fmt.Println(bytes)
	runes := []rune(new)
	fmt.Println(runes)

}
