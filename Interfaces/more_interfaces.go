package main

import "fmt"

func main(){

	var i interface{}  // Empty interfaces
	i = "Hello"	// Maybe use this when I am not sure what data type I may get ?

	j := i.(string)	// Assertion,  it checks for underlying data structure, assertion only works on interface, while conversion works all the way 
	k, ok := i.(int)
	fmt.Println(j, k, ok) // Hello 0 false

	// If you do not use k, (ok) idiom this can cause panic like in next line

	m := i.(int) // panic: interface conversion: interface {} is string, not int
	fmt.Println(m)
}
