package main

import (
	"fmt"
)

/* Zero value for an interface is nil
This nil is different from nils of other data types like
Maps, slices, runes etc. The underlying concrete value nil 
can be different in case of interface and other data types. 
*/

func reallyNil() error{
	var e error
	fmt.Println("e is nil", e == nil)
	return e
}

// MyError struct
type MyError struct{
	A int
	B int
	message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error %s", me.A, me.B, me.message)
}

func notReallyNil() error{
	var me *MyError
	fmt.Println("me is nil", me == nil)
	return me
}

func main() {
	e := reallyNil()
	me := notReallyNil()
	fmt.Println("In main, e is nil", e == nil)
	fmt.Println("In main, me is nil", me == nil) // magic happens here
	// even tho the value returned is nil, but this nil is coming from nil of *MyError, which is not true nil for interface.
}
