package main

import (
	"sync"
	"fmt"
)

func runMe(){
	fmt.Println("Hello from Go routine!")
}
func main(){
	// We use WaitGroup instead of sleep
	var wg sync.WaitGroup // We declare wg of type wait group
	wg.Add(1) // add 1 go routine
	go func() { // yep, A closure
		runMe()
		wg.Done()
	}()

	wg.Wait()
}
