package main

import "fmt"

func main() {
	var firstName string
	firstName = "Rahiyan"

	lastName := "Safin"

	fmt.Println(firstName)
	fmt.Println(lastName)

	var currentYear int = 2023
	birthYear := 1998

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear = currentYear + 1

	fmt.Println(currentYear)

	age = currentYear - birthYear

	fmt.Println(age)
}
