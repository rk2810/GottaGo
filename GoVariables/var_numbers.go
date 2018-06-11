package main

import "fmt"

func main() {
	/* Variables are not as easy as they seem in Go
	There are multiple ways and each way is unique
	Also, Every variable that is declared must be used.*/


	var i byte = 10 // Change byte to int and you will see a warning, cuz literal assigned to i on right is an int.
	fmt.Println(i)
	var j = 20 // An integer as depicted on right hand side
	fmt.Println(j)

	/* There is something called zero value of a data type, other languages like C show them as undefined, yet in Go
	It assigns zero value to declared but unassigned variables*/

	var k int
	fmt.Println(k) // Should print 0 as for int, Zero value is 0
	k = 30
	fmt.Println(k)

	/* There is another way of making variables */

	// var l int
	l := 30 
	/* beware := means declaration and not assignment so already declared variable will throw error
	Uncomment line 26 and you will see.*/
	fmt.Println(l)
}
