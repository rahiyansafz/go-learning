package main

import "fmt"

func main() {
	luckyNumber := 7
	evenMoreLuckyNumber := luckyNumber + 2

	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = evenMoreLuckyNumber + 2

	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)
}
