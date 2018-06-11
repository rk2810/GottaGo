package main
import (
	"fmt"
)

func main() {
	var s, s1 string
	s = `Hey there, `
	s1 = "What's up !"
	var s3 rune = 127757 // A globe,  rune is just an int32 char, we basicaly use rune to handle unicode shits properly
	fmt.Println(s + s1 + string(s3))
	fmt.Println(len(s)+len(s1)+len(string(s3)))

	// Arrays
	var x [4]int 
	x[0] = 1
	// x[1] = 2  Zero value of int will kick in
	x[2] = 3
	x[3] = 4
	fmt.Println(x)


	/* There is a huge shitty shit in Go when using arrays.
	The arrays in Go are fixed size, unlike python, you can't have variable length of arrays
	Infact, you'll face issues in following operations:
	var x [4]int
	var z [5]int = x
	>>> cannot use x (type [4]int) as type [5]int in assignment

	WE USE SLICE FOR SUCH STUFF..  **THE MORE YOU KNOW**
	*/

}
