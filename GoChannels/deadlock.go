package main

import (
	"fmt"
)

func main(){
	in := make(chan int) // Just pass 2 here and it will work just fine
	out := make(chan int)

	go func(){
		for {
			i := <-in
			out <- i*2
		}
	}()

	in <- 1
	o1 := <- out
	in <- 2
	o2 := <- out

	// It's pretty clear that program doesnt know how many values its gonna get and hence the deadlock

	fmt.Println(o1, o2)
}
