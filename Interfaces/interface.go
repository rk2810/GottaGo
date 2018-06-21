package main

import (
	"fmt"
)

// Interfaces are list of methods, which are yet to be implemented
// Interfaces wrap stuff from their underlying data type

type tester interface{ // Declare an interface
	test(int) bool	  // In interface, we define the method signature
				     // i.e Method name, argument and return type.
}


func runTest(i int, tests []tester) bool{ // Function runTest expects an int i and a slice of type tester and returns bool
	result := true
	for _, x := range tests{
		result = result && x.test(i) // calling the interface to check if given int argument complies 
		fmt.Println("x is ", x, "result is", result)
	}
	return result // A bool
}

type rangeTest struct {  // A structure for 2 ints
	min int
	max int
}

func (rt rangeTest) test(i int) bool { // We declare a method that takes rt rangeTest and retruns a bool
	return rt.min <= i && i <= rt.max
}

type divTest int  // A type based on int

func (dt divTest) test(i int) bool {  // A method called test() on divTest and returns bool 
	fmt.Println(i)
	fmt.Println(dt)
	return i%int(dt) == 0  // returns bool for, if int is divisible by dt
}

func main() {	
	result := runTest(10, []tester{
		rangeTest{min:5, max:20}, // Supplied max and min for struct
		divTest(5), // div test
	})
	fmt.Println(result)
}