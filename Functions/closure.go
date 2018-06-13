package main

import "fmt"

/*A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

func makeAdder(b int) func(int) int { // This function accepts an int, yet, returns a function
	return func(a int) int{
		fmt.Println("a is",a)
		fmt.Println("b is",b)
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int { // Now this mofo expects a function in args and returns a function as well
	return func( b int) int{
		fmt.Println("b is",b)
		z := f(b)
		fmt.Println("z is",z)
		return z * 2
	}
}

func main() {
		// Call makAdder function
		x := makeAdder(2)	// Assign function to a variable and pass int argument
		fmt.Println(x(3))	// The function then returns a function, which is also expecting an int argument

		y := makeDoubler(x)
		fmt.Println(y(5))

}