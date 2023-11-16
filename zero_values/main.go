package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		value       int     // this is automatically initialized to zero(0)
		name        string  // this is automatically initialized to empty string("")
		salary      float64 // this is automatically initialized to zero(0)
		isDeveloper bool    // this is automatically initialized to false
	)
	fmt.Println(value, name, salary, isDeveloper)
	fmt.Printf("the value of 5 raise to power 2 is : +%f \n", math.Pow(5, 2))
	fmt.Printf("he says and I quote: \"I'm a Developer\"")
}
