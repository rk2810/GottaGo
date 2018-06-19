package main

import "fmt"

// Foo is a struct with A int and B string
type Foo struct {
	A int
	B string
}

// Concept of methods, methods vs. functions
// Go does not have classes, yet we can use structs and pass them in functions


// fieldCount method returns an int 2
func (f Foo) fieldCount() int{
	return 2
}

// String method has a reciever of type Foo named f
func (f Foo) String() string {
	return fmt.Sprintf("%d Fields: A : %d and B: %s", f.fieldCount(), f.A, f.B)
}

// Double method has a reciever of type pointer to Foo named f
func (f *Foo) Double() {
	f.A = f.A * 2
}

// Bar is a struct with a bool and type Foo struct
type Bar struct {
	C bool
	Foo
}

// Methods operating on Bar

func (b Bar) String() string{
	return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
}

func (b Bar) fieldCount() int {
	return 3
}


type myInt int

func (mi myInt) isEven() bool{
	return mi%2==0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

func main() {
	//  Method is a function with a special reciever argument
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())

	b := Bar{
		C: true,
		Foo: f,
	}

	// Surprise Surprise !
	fmt.Println(b.String()) 
	b.Double()
	fmt.Println(b)
	
	/*Output >
	2 Fields: A : 20 and B: Hello and C: true
	2 Fields: A : 40 and B: Hello and C: true

	There is no inheritence or method override in Golang, So what happened is,
	b.String() method has b.Foo.String(), Now it goes to String() for Foo and String() for Foo calls fieldCount()
	of Foo, and it has no idea if a method of same name exists for Bar.

	> No method overriding
	> No virtual method dispatch

	*/
	 
	z := myInt(10)  // this is type cast
	fmt.Println(z)
	fmt.Println(z.isEven())
	z.Double()
	fmt.Println(z)


	}
