package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // Add a wait gorup
		go func(LocalI int){ // If we dont do this, you will just see bunch of 10s in output
			fmt.Println(LocalI)
			wg.Done()
		}(i) // Pass i in closure func
	}
	wg.Wait() // Wait till all routines complete
}
