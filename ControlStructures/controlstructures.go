package main
import "fmt"
func main() {
	// If-Else
	fmt.Println("Demo of If-else:")
	fmt.Println()

	var a = 10
	fmt.Println("a is", a)
	if a := 4; a > 5 { // yea, you witnessed a reassignment during an if statement
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is smaller than 5. It's", a)
	}
	fmt.Println("Value of a is back to", a)
	fmt.Println()
	/* Right, the value is back to 10
	since the reassignment scope was inside if-else block only.
	*/

	//For loops
	fmt.Println("Demo of for loop:")
	fmt.Println()
	for i := 0; i < 10;i++ { // remember, NO PARANTHESIS and mind the curly braces.
		if i > 5{
			break // Just using break, can use continue
		} else if i == 2 { // example of else if
			continue // 2 will be just skipped
		}
		fmt.Println(i)
	}

	// While loops (Go has no while loop) so we use for loop like a while loop without increment
	fmt.Println("Demo for While loop")
	fmt.Println()
	i := 0
	for i < 10{
		fmt.Println(i)
		i++
	}

	// Infinite loops
	fmt.Println("Demo for infinite loops")
	fmt.Println()
	j := 0
	for{
		fmt.Println(j)
		j++
		if j > 9 { // Of course I wont run an infinite loop on purpose.
		break
		}
	}

	// Range
	fmt.Println("Range can be used for iteration through some Go built in types")
	fmt.Println()
	s := "Hello World!"
	for k, v := range s { // here k is Offset or index, and v is Rune, S is string as defined
		fmt.Println(k, string(v), v)
	}
}
