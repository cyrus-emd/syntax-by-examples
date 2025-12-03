package main

import (
	"fmt"
)

// Package level:
var userCountry = "Iran"

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {
	// Function level:
	var userName, userAge = "Zahra Emadi", 23
	var isMale bool
	fmt.Println(userName, "comes from", userCountry+".")
	fmt.Println(userName, "is", userAge, "years old.")
	fmt.Println(userName, "is male ->", isMale)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type:
	userCity, siblingsCount := "Gorgan", 1
	fmt.Println(userName, "lives in", userCity+".")
	fmt.Println(userName, "has", siblingsCount, "brother.")
}
