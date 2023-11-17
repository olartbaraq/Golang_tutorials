package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Of course, we all know functions and all.
// My interest in this class is to be able to return multiple values
// at the same time syntactically

// func checkErr(file *os.File, err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("File created successfully")
// }

//
//
//

// Variadic functions
// Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that parameter.
// If the function takes parameters of different types, only the last parameter of a function can be variadic.

// creating a variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a) // => []int, slice of int
	fmt.Printf("%#v\n", a)
}

// variadicfunction that modifies one of the arguments passed.
// func f2(a ...int) {
// 	a[0] = 50
// }

// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

// DEFER STATEMENT.
// defer statement postpone execution of a function until the surrounding functions are executed
func foo() {
	fmt.Println("Hello Foo!")
}
func bar() {
	fmt.Println("Hello bar!")
}
func fooBar() {
	fmt.Println("Hello FooBar!")
}

// ANONYMOUS FUNCTIONS
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

// exercise
func myFunc(n string) int {
	num1 := strings.Repeat(n, 2)
	num2 := strings.Repeat(n, 3)

	nInt, err1 := strconv.Atoi(n)
	num1Int, err2 := strconv.Atoi(num1)
	num2Int, err := strconv.Atoi(num2)
	if err1 != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return nInt + num1Int + num2Int
}

func main() {
	// calling variadic functions
	// a variadic function can be invoked with zero or more arguments
	f1(1, 2, 3, 4)

	f1() // a is: []int(nil)

	// an example of a well-known variadic function is append() built-in function.
	// it appends items to an existing slice and returns back the same slice.
	nums := []int{1, 2}
	nums = append(nums, 3, 4)
	_ = nums

	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p) // -> 17 100

	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart

	fmt.Println(strings.Repeat(("#"), 20))

	defer foo()    // 3
	bar()          // 1
	defer fooBar() // 2

	//
	//
	// ANONYMOUS FUNCTIONS
	// they are functions that doesn't contain any name and are declared inline using function literal
	// they are use to form closures.

	func(msg string) {
		fmt.Println(msg)
	}("This is an anonymous function")

	a := increment(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(strings.Repeat(("#"), 20))

	fmt.Println(myFunc("9"))

}
