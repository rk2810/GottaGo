package main

import (
	"fmt"
	"os"
)

func main() {
	in := make(chan string) // chan keyword to create channel
	out := make(chan string)
	go func(){
		name := <-in // Take value from channel
		out <- fmt.Sprintf("Hello, " + name + ".") // Write value to channel
	}()
	in <- os.Args[1] // Take input from user
	close(in)
	message := <-out
	fmt.Println(message)
}
