package main

import (
	"fmt"
)

// Main Function
func main() {
	// Variable declaration/use
	//no unused variable allowed

	var name string = " Muhammad "// type inferred and shorthand  operator(:=)
	age := 414
	fmt.Printf(" This is my first name is %s\n and I am %d\n years old,", name, age)

	var city string // variable declaration without initialization
	city = "Great London"
	fmt.Printf(" and I live in %s.\n", city)

	var country, continent string = "England", "Europe"
	fmt.Printf("My country is %s\n and continent is %s\n", country, continent)

var (
	isEmployed bool = true
	salary int = 52000
	position string = "Data scientist")
	fmt.Printf("isEmployed : %t\n this is my Yearly %d\n salary and my position is %s\n", isEmployed, salary, position)

	// zero Values, compiler will assign the values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("default int: %d\n default string: %s\n default float: %f\n default bool: %t\n", defaultInt,defaultString,defaultFloat,defaultBool)

// Const Values
	const pi = 3.14 //const values can be left unuse

// declaring const variable with same type
	const (
		Monday = 1
		Tuesday = 2
		Wednesday = 3
	)
	fmt.Printf("  Monday: %d\n  Tuesday: %d\n  Wednesday: %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 25
	const untypedAge = 25
	fmt.Println(typedAge == untypedAge)

// Enum like Structure
	const(
		Jan = iota +1
		Feb
		March
		April
		May
	)
	fmt.Printf(" Jan: %d\n Feb: %d\n March: %d\n April: %d\n May: %d\n", Jan, Feb, March, April,May)

//add Function
result := add(625, 5483)
fmt.Printf("This is the result %d ", result)

//Multiple Return Types
sum, product := calculateSumAndProduct(2 ,5)
fmt.Printf("The sum is %d\n and the product is %d\n", sum, product)
}

//add Function
func add(a int, b int) int {
	return(a+b)
}
// Multiple Return Types
func calculateSumAndProduct(a,b int)(int, int){
	return a+b,a*b
}
