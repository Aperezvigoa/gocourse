package basics

import "fmt"

var middleName string = "Calonge"

func main() {
	// var age int
	// var firstCatName string = "Cheddar"
	// var secondCatName string = "Bigotitos"

	// count := 10
	// lastName := "Albizu"

	// Default Values
	// Numeric Types: 0
	// Boolean:  false
	// String type: ""
	// Pointers, slices, maps, functions and structs: nil

	// ------------ SCOPE
	middleName := "Perez-Vigo"
	fmt.Println(middleName)
}

func printName() {
	firstName := "Lilui"
	fmt.Println(firstName)
}
