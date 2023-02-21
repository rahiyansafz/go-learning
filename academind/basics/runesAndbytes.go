package main

import "fmt"

func main() {
	var runeVal rune = '$'
	fmt.Println(runeVal)
	fmt.Println(string(runeVal))

	byteVal := '$'
	fmt.Println(byteVal)
	fmt.Println(string(byteVal))
}
