package main

import "fmt"

/* Select is like cops in go routines, they check feasibility of channel read/write
It looks like a switch statement, so basically it's like a control structure.
*/

func main(){
	in := make(chan int) // redefined in and out for unbuffered
	out := make(chan int)

	select{ // Select statement picks a case randomly to check feasibility
	case in <- 2:
		fmt.Println("Wrote 2 to in.")
	case v := <-out:
		fmt.Println("Read", v, "from out.")
	default:
		fmt.Println("Nothing worked.")
	}
	/* 
	
	[rohank@Cerebrum GottaGO]$ go run GoSelect/select.go
	Nothing worked.

	*/
}
