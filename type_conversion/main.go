package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, b, y = 4, 7.8, 5.6
	//var y = 5.6

	fmt.Println(float64(x) + y + b) // mismatch type

	// but i can declare a variable which can  result to operation of different types
	//e.g
	result := 5.5 * 7
	fmt.Println(result) // this will work without compile time error because result is being
	//inferred to float64 type

	//CONVERTING INT TO STRING AND VICE VERSA
	a := 90
	s := string(a) // the string conversion type return a rune i.e
	// a character
	j := fmt.Sprintln(a) // to conver a value like an int or float to a proper string , use fmt.Sprint function

	fmt.Printf("The ascii value of %q is %v \n", a, s)
	fmt.Println(j)

	//CONVERTING STRING TO NUMBERS (INT OR FLOAT)
	s1 := "3.456"

	f1, err := strconv.ParseFloat(s1, 64)

	i, err2 := strconv.Atoi("34465")

	_, _ = err, err2

	//si := strconv.Itoa(34)
	// or
	si := fmt.Sprintln(34)

	fmt.Println(f1)
	fmt.Println(i)
	fmt.Println(si)

}
