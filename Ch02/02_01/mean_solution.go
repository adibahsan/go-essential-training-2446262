// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 1
	y = 2

	xFloat := 1.00
	yFloat := 2.00

	/*
		var x float64
		var y float64

		x = 1.0
		y = 2.0
	*/

	/*
		x := 1.0
		y := 2.0
	*/

	/*
		x, y := 1.0, 2.0
	*/

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	meanFloat := (xFloat + yFloat) / 2

	fmt.Printf("result: %v, type of %T\n", mean, mean)
	fmt.Printf("result: %v, type of %T\n", meanFloat, meanFloat)
}
