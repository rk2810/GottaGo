package main

import "fmt"

/* Select is like cops in go routines, they check feasibility of channel read/write
It looks like a switch statement, so basically it's like a control structure.
*/

func main(){
	in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	select{ // Select statement picks a case randomly to check feasibility
	case in <- 2:
		fmt.Println("Wrote 2 to in.")
	case v := <-out:
		fmt.Println("Read", v, "from out.")
	}
	/* So when you run this, you can have a random output.
	
	[rohank@Cerebrum GottaGO]$ go run GoSelect/select.go
	Wrote 2 to in.
	[rohank@Cerebrum GottaGO]$ go run GoSelect/select.go
	Read 1 from out.

	*/
}
