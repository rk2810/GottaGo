package main

import (
	"strconv"
	"os"
	"fmt"
)

/* Golang does not have error handling like C, but it does have a way to deal with it.
Since error here is an interface, we can create our own error handling mechanism.
*/

type myError struct{ // Custom error struct
	A int
	B int
	message string
}

func (me myError) Error() string{ // Error method
	return fmt.Sprintf("values %d and %d produced error: %s", me.A, me.B, me.message)
}

func divAndMod (a int, b int) (int, int, error) {
	if b == 0{
		return 0, 0, &myError{A:a, B:b, message:"cannot divide by zero."} // Instanciated myError 
	}
	return a / b, a % b, nil
}

func main(){
	if len(os.Args) < 3{
		fmt.Println("Expected two input params.")
		os.Exit(1)
	}
	a ,err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	b ,err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	div, mod, err := divAndMod(int(a), int(b))
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d divided by %d is %d, the remainder is %d\n", a, b, div, mod)
}
