package main

import "fmt"

func main() {
	const PI = 3.14
	// PI := 3.14
	radius := 5

	circuleCircumference := 2 * PI * float64(radius)
	// circuleCircumference := 2 * int(PI) * radius

	fmt.Printf("For a radius of %v, the circle circumference is %.2f.", radius, circuleCircumference)
}
