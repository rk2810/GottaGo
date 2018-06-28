package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	var wg sync.WaitGroup
	totalStart := time.Now()
	for i := 0; i < 1000; i++ { // We are gonna loop 1000 time, wiz. 1000 read/writes
		start := time.Now()
		wg.Add(1)
		go func(in int) {
			time.Sleep(1 * time.Second)
			out := 2 * in
			fmt.Println("Got", out, "for", in, "after", time.Now().Sub(start))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Total time:", time.Now().Sub(totalStart))
}
/* There can be times when we get too many requests and dont want to swamp our resource, we can create a worker pool to
   for such instances. This is called back pressure.
*/