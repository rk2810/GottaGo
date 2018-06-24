package main

import (
	"time"
	"fmt"
)

func runMe(){
	fmt.Println("Hello from Go routine!")
}
func main(){
	go runMe()
	/* This is just an example, should never do this irl, cuz we never know how long we
	   need to wait for go routine to finish */
	time.Sleep(1 * time.Second) // sleep for a second
}
