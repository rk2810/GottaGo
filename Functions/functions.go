package main

import (
	"fmt"
)

// Go does not have features like functional overloading
// GO uses Call by value concept** and not call by reference. Hence the originl values are persistent. 

func add(a int, b int) {
	fmt.Println(a + b)
}

func division(a int, b int) (int, int) { // input parameters and return types
	return a / b, a % b
}

/* Time to step up our game and pass function as a param in functions
I will call them in main
*/

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int{
	return a + 2
}

func printOperation(a int, f func(int)int){
	fmt.Println(f(a))
}

func main() {
	add(1,2)

	div, remainder := division(10, 2)
	fmt.Println(div, remainder)

	div, _ = division(5, 5) // In Go, you just can't leave a variable, hence the underscore.
	fmt.Println(div)

	_, remainder = division(1, 2)
	fmt.Println(remainder)

	// defining function inside another function in Go.

	newFunc := func (a int) int { //if you go by regular method you will face compilation error.
		return a
	}
	fmt.Println(newFunc(15))
	
	// The calling 
	printOperation(5, addTwo)
	printOperation(10, addOne)
}
