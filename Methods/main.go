package main

import (
	"fmt"
	"strings"
	"time"
)

type names []string

func (n names) print() {
	for i, v := range n {
		fmt.Println(i, v)
	}
}

// here to defined a method to be called as a receiver function, we pass the underlying type as a parameter
// before the name of the function is written.

type car struct {
	model string
	price int
}

func changeCar(c car, newModel string, newPrice int) {
	c.model = newModel
	c.price = newPrice
}

func (c *car) changeCarRe(newModel string, newPrice int) {
	c.model = newModel
	c.price = newPrice
}

// func changeCarPointer(c *car, newModel string, newPrice int) {
// 	c.model = newModel
// 	c.price = newPrice
// }

//
//
//NB: we can only create receiver for value types and not pointer types.

// type value *int

// func (v value) money() { // invalid in golang
// 	fmt.Println("Just printing")
// }

func main() {

	// Since there is no class in Golang, we can define Methods on Predefined types

	const day = 24 * time.Hour
	fmt.Printf("The type of day is: %T and its value is :%v\n", day, day) //time.Duration

	// The type of the const day here is sterning from the time package, which makes day
	// the type of time.Duration.

	minutes := day.Minutes()
	seconds := day.Seconds()
	fmt.Println(minutes, seconds)

	fmt.Println(strings.Repeat("-", 20))

	// since we were able to have a type time.Duration applied on a primitive type int
	// we can then proceed to use methods on that type like the minutes() and seconds()
	// These are called RECEIVER FUNCTIONS.

	//
	//
	// ** OUR OWN NAMED TYPE AND METHOD **//
	friends := names{"Mubby", "Qudus"}
	// since names is a defined type of slice string([]string), friends is a variable of the type name and
	// can proceed to call the print() on the variable friends.

	friends.print() // I like this one better; easier to remember. just like self.print()
	// OR IT CAN ALSO BE CALLED LIKE
	names.print(friends) // like class.method()

	fmt.Println(strings.Repeat("-", 20))

	//
	//
	//
	//It's idiomatic to convert the type of an expression to accept a method.

	timeInt := 12353456324546
	fmt.Println(time.Duration(timeInt))

	//
	//
	//

	//** METHODS WITH A POINTER RECEIVER **//

	toyota := car{model: "camry", price: 2000}
	changeCar(toyota, "corolla", 1500)
	fmt.Println(toyota) // no changes because this function worked on a copy.

	toyota.changeCarRe("corolla", 2500)
	fmt.Println(toyota)

	// OR //
	(&toyota).changeCarRe("Pathfinder", 25000)
	fmt.Println(toyota)

	// changeCarPointer(&toyota, "Pathfinder", 2500)
	// fmt.Println(toyota)

}
