package main

import "fmt"

func main() {
	var isMale bool // zero value: false
	var isYoung = true
	fmt.Printf("%v: %T\n", isMale, isMale)
	fmt.Printf("%v: %T\n", isYoung, isYoung)

	var firstName string // zero value: "" (empty string)
	var lastName = "Amirhossein Emadi"
	fmt.Printf("%v: %T\n", firstName, firstName)
	fmt.Printf("%v: %T\n", lastName, lastName)

	var number1 int // zero value: 0
	var number2 = 25
	fmt.Printf("%v: %T\n", number1, number1)
	fmt.Printf("%v: %T\n", number2, number2)

	var number3 float64 // zero value: 0
	var number4 = 2.5
	fmt.Printf("%v: %T\n", number3, number3)
	fmt.Printf("%v: %T\n", number4, number4)

	// rune = int32
	var rune1 rune // zero value: 0
	var rune2 = 'C'
	fmt.Printf("%v: %T\n", rune1, rune1)
	fmt.Printf("%v: %T\n", rune2, rune2)

	var cpx1 complex128 // zero value: (0+0i)
	var cpx2 = 1278i
	fmt.Printf("%v: %T\n", cpx1, cpx1)
	fmt.Printf("%v: %T\n", cpx2, cpx2)
}
