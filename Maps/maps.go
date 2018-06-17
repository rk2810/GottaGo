package main

import (
	"fmt"
)

func main() {
	// Just like slices, there are 3 ways to initialize a map, Map is just like a dict in python BTW
	
	// 1. Make method

	m := make(map[string]int)
	m["x"] = 300
	m["y"] = 200
	m["z"] = 100

	fmt.Println(m) // the order of key value is totally random when printed.

	// range convention for maps to check if key exists
	if v, ok := m["x"]; ok { // v is value and ok is a conventional thingy, it returns true or false
		fmt.Println("x in m is ", v)
		fmt.Println(ok)
	}
	if v, ok := m["a"]; ok {
		fmt.Println("a in m is ", v)
	} else{
		fmt.Println("Not found.")
	}

	// 2. IDK what to call this method of declaration

	m1 := map[string]int {
		"a" : 100,
		"b" : 200,
		"c" : 300, // take note, even the last element has ,
	}
	fmt.Println(m1)

	// deleting element by key
	delete(m1, "b")
	fmt.Println(m1)

	// 3. The var method
	
	var m2 map[string]int
	fmt.Println(m2)
	m2 = m
	fmt.Println(m2)

	modMap(m2) // This func is just to show the referential thing we have in maps
	fmt.Println("m2 is") // Since m2 = m from lone 47
	fmt.Println(m2) // making changes in m2 reflects on m, This must be kept in mind when passing maps in functions
	fmt.Println("m is")
	fmt.Println(m)

	// I want to try adding a value to a nil map

	// var m3 map[string]int
	// modMap(m3)
	// fmt.Println(m3)
	
	// panic: assignment to entry in nil map 

}

func modMap(m map[string]int) {
	m["n"] = 69
}