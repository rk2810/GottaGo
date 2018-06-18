package main

import (
	"encoding/json"
	"fmt"
)

// Person struct takes name and age of a person.
type Person struct {
	Name string `json:"name"`  // This is a string literal defining a json, This is called a struct tag
	Age int     `json:"age"`
}

func main() {

	bob := `{"name":"bob", "age":30}`
	var b Person
	json.Unmarshal([]byte(bob), &b) // A byte slice of bob and pointer to b
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}