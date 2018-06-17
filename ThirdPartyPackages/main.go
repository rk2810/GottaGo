package main


import (
	"fmt"

	"github.com/rk2810/GottaGO/ThirdPartyPackages/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Hi there")) // English
	fmt.Println(mapper.Greet("Comment allez vous?")) // French
	fmt.Println(mapper.Greet("Wie geht es Ihnen?")) // German
}
