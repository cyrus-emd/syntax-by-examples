package main

import (
	"fmt"
	"math"
)

// In Go, a name is exported if it begins with a capital letter.
// When importing a package, you can refer only to its exported names.

// userId an unexported name.
const userId = 1234

// UserName an exported name.
const UserName = "Amirhossein Emadi"

func main() {
	// Pi is an exported name from math package:
	fmt.Println(math.Pi)
	fmt.Println(UserName+":", userId)
}
