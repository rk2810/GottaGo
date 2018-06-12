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
func main() {
	add(1,2)

	div, remainder := division(10, 2)
	fmt.Println(div, remainder)

	div, _ = division(5, 5) // In Go, you just can't leave a variable, hence the underscore.
	fmt.Println(div)

	_, remainder = division(1, 2)
	fmt.Println(remainder)
}
