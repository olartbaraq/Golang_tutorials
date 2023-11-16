package main

import (
	"fmt"
	//"math"
)

//"fmt"

func main() {
	// rule 1 : you can't chnage a value of a const
	const temp int = 300
	//temp = 50  // not possible

	// rule 2
	//const power float64 = math.Pow(5, 2)
	// cannot initialze a const at runtime	 math.Pow belongs to a runtime funtion

	// rule 3
	//cannot use a variable to initialize a constant
	// age := 5
	// const myAge = age

	//rule 4
	//len funtion (len()) with a literal string can be used to initialize a constant because the len
	// function belongs to a compile-time constant
	//const myNum = len("myAge")

	// untyped constants

	const timeInSec = 3600      // This is an unyped constant
	const timeInMinute int = 60 // This is a typed constant

	fmt.Printf("The type of %q is %T and their value is %v \n", ("2 * 5.5"), (2 * 5.5), (2 * 5.5)) // 11 no error because the values are both untyped
}
