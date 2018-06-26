package main

import (
	"fmt"
)

func main() {
	out := make(chan int, 10) // A room for 10 ints in a channel
	for i := 0; i < 10; i++ { // A loop for i
		go func(LocalI int){
			out <- LocalI * 2 // write value to out channel
		}(i)   // A go routine that takes i and  doubles its value
	}
	var result []int
	for i := 0; i < 10; i++ { // same loop to get 10 values from go routine
		val := <-out
		result = append(result, val) // append them to result slice
	}
	fmt.Println(result)
}
