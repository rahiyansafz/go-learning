package main

import "fmt"

func main() {
	firstName := "Rahiyan"
	lastName := "Safin"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 26
	occupation := "Software Engineer"

	fmt.Printf("Hi, My name is %v and I am %v years old and I am a %v.", fullName, age, occupation)

	fmt.Println(" ")
	fmt.Printf("Firstname Type is %T", firstName)
	fmt.Println(" ")
	fmt.Printf("Lastname Type is %T", lastName)
	fmt.Println(" ")
	fmt.Printf("Age Type is %T", age)
	fmt.Println(" ")
	fmt.Printf("Occupation Type is %T", occupation)
}
