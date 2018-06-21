package main

import "fmt"

/* Interface can math methods with any types, not just structs
Go let's you make a func implement an interface, by defining a
function type and method on that function type.
*/

type tester interface {
	test(int) bool
}

func runTest(i int, test []tester) bool {
	result := true
	for _, x := range test {
		result = result && x.test(i)
	}
	return result
}

type testerFunc func(int) bool // A type testerFunc of type func and it takes int as param and returns bool

func (tf testerFunc) test(i int) bool { //A method on testerFunc that calls a function test
	return tf(i)
}

func main() {

	result := runTest(10, []tester{
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result)
}
