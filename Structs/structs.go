package main
import "fmt"

// Foo is a struct here, We dont have inheritance in Go, and hence structs are different than objects
type Foo struct {
	A int
	b string
}
func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	fmt.Println(g.b)
	f.b = "Hey"
	fmt.Println(f)
}