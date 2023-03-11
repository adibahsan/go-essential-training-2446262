/*
Print the numbers between 1 to 20, however
- If the number is divisible by 3, print "fizz" instead
- If the number is divisible by 5, print "buzz" instead
- If the number is divisible by 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter A Number")
	var a int

	fmt.Scan(&a)

	switch {
	case a%5 == 0 && a%3 == 0:
		{
			fmt.Println("fizzbuzz")
		}
	case a%5 == 0:
		{
			fmt.Println("buzz")
		}
	case a%3 == 0:
		{
			fmt.Println("fizz")
		}
	default:
		fmt.Println(a)
	}

}
