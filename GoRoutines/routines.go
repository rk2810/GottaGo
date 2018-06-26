package main

import (
	"sync"
	"fmt"
)

func runMe(name string){
	fmt.Println("Hello to,", name, "from Go routine!")
}
func main(){
	// We use WaitGroup instead of sleep
	var wg sync.WaitGroup // We declare wg of type wait group
	wg.Add(1) // add 1 go routine
	go func(name string) { // yep, A closure
		runMe(name)
		wg.Done()
	}("Rohan")

	wg.Wait()
}
