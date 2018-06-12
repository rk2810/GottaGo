package main

import (
	"os"
	"fmt"
)

func main() {
	word := os.Args[1]
	/* Regular if else control

	if word == "hello"{
		fmt.Println("Hey yourself!")
	} else if word == "bye"{
		fmt.Println("Goodbye!")
	} else if word == "Greetings"{
		fmt.Println("Hola, Come esta!")
	} else {
		fmt.Println("Don't know...")
	}
	
	We can deploy switch case here:
	*/
	
	switch word {
	case "hello":
		fmt.Println("Hey yourself !")
	case "bye", "later": // Yea, I can do that
		fmt.Println("Goodbye.")
	case "greetings":
		fmt.Println("Salutations.")
		fallthrough // Makes no sense here, but just to demonstrate 
		// Unlike C/C++/Java, we do not break each statement, we use fallthrough instead.
	default:
		fmt.Println("I don't understand.")
	}
}
