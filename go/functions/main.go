package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("3 + 8 =", addTwoNumbers(3, 8))
	fmt.Println("3 - 8 =", minusTwoNumbers(3, 8))
	checkZero, result := divideTwoNumbers(13, 7)
	if checkZero {
		fmt.Println("13 / 7 =", result)
	}
	multiply, power := multiplyAndPowerTwoNumbers(2, 3)
	fmt.Println("2 * 3 =", multiply)
	fmt.Println("2 ** 3 =", power)
}

// Notice that the type comes after the variable name:
func addTwoNumbers(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last:
func minusTwoNumbers(x, y int) int {
	return x - y
}

// A function can return any number of results:
func divideTwoNumbers(x, y float64) (bool, float64) {
	if y == 0 {
		return false, 0
	}
	return true, x / y
}

// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions. They can harm readability in longer functions.
func multiplyAndPowerTwoNumbers(x, y float64) (multiply, power float64) {
	multiply = x * y
	power = math.Pow(x, y)
	return
}
