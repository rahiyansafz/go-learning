package main

import "fmt"

func output(name string, value interface{}) {
	fmt.Printf("%s: %v\n", name, value)
}
