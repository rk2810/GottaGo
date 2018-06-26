package main

import "fmt"

func main(){
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++{
		in <- i
	}
	close(in) // closing a channel doesnt wipe the values, they are available to be read.
// if channel is empty, it returns zero value for channel type, but how do you differentiate a real zero from this Zero value?

	go func(){
		for{
			i, ok := <- in // Here's how, i, ok idiom
			if !ok{
				close(out) // close channel if not ok
				break
			}
			out <- i*2
		}
	}()
// Never write to a closed channel, or close a channel twice, this can panic the program.
	for v := range out{
		fmt.Println(v)
	}
}
