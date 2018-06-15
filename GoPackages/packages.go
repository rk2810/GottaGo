package main

import (
	"unicode/utf8"
	"fmt"
	"strings"
	"time"
	"math"
	"math/rand"
)

func rot13(in rune) rune {
	// Let's try implementing rot13 encryption
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}
	return in

}

func main() {
	
	// Some functions from strings package
	string1 := "Hello 🌍!"
	encrypted := strings.Map(rot13, string1) // Map return a copy of string
	fmt.Println("Encrypted string >", encrypted)
	decrypted := strings.Map(rot13, encrypted)
	fmt.Println("Decryted string >", decrypted)
	// s3 := strings.Map(rot13, "hello")
	// fmt.Println(s3)

	// Now let's try to get length of strings.
	fmt.Println("Length of string without considering runes :", len(string1)) // 11, crazy right ?

	fmt.Println("Length of string without considering runes :", utf8.RuneCountInString(string1)) // Sanity.

	timenow := time.Now() // time.Time type object
	fmt.Println("Current time is ", timenow)
	nanos := timenow.UnixNano() // Epoch time int64
	fmt.Println("Epoch time is", nanos)

	// Using math package
	fmt.Println("Value of cosine of Pi is", math.Cos(math.Pi))

	var x float64 = 10
	var y float64 = 20
	fmt.Println("Max number is ", math.Max(x, y))
	fmt.Println("Min number is ", math.Min(x, y))

	rand.Seed(nanos) // Initialize seed value, using epoch time here
	//  get some pseudo random numbers
	a := rand.Intn(100)
	b := rand.Intn(100)
	fmt.Println(a, b)

}
