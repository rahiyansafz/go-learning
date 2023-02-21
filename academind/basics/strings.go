package main

import "fmt"

func main() {
	var greetingWithVar string = "Hello World!"

	greetingWithoutVar := "Hello World!"

	var greetingWithoutValue string
	greetingWithoutValue = "Hello World!"

	fmt.Println("greetingWithVar: ", greetingWithVar)
	fmt.Println("greetingWithoutVar: ", greetingWithoutVar)
	fmt.Println("greetingWithoutValue: ", greetingWithoutValue)
}
