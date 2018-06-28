package main

import (
	"time"
	"fmt"
)

/*
Go manages its resources but not for channels, If you do not close channels and clear buffer,
You will end up swamping the resources. This can be done with done cahnnel.
*/

func multiples(i int) (chan int, chan struct{}){	// takes integer argument and returns a channel
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func(){ 
		for{ // an infinite for loop that takes current value and doubles and increments it.
			select{
			case out <- curVal*i:
				curVal++
			case <- done:
				fmt.Println("Closing go routine.")
				return
			}
		}
	}()
	return out, done
}

func main(){
	twoCh, done := multiples(2) // twoCh calls func multiples and opens a channel
	for v := range twoCh{
		if v > 20{ // value of v reaches 20 and loop breaks
			break
		}
		fmt.Println(v)
	}
	close(done)
	time.Sleep(2 * time.Second)
}
