package main

import (
	"sync"
	"fmt"
	"time"
)

// Using back pressure, we can decide the worker pool and limit the number of requests.
func main() {
	worker := 1000
	pool  := make(chan func(int) int, worker) // here we defined a pool of functions
	for i := 0; i < worker; i++{
		pool <- func(in int) int{
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < 1000; i++{
		start := time.Now()
		select{ // a select
		case f := <- pool: // if we have a worker
			fmt.Println("Processing", i)
			count++
			wg.Add(1)
			go func(in int){
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)
		default: // We do not have a worker to spare
			fmt.Println("Rejecting request", i, "too busy")

		}
	}
	wg.Wait()
	fmt.Println("Total processed:", count)
	fmt.Println("Total time:", time.Now().Sub(totalStart))
}
/* There can be times when we get too many requests and dont want to swamp our resource, we can create a worker pool to
   for such instances. This is called back pressure.
*/