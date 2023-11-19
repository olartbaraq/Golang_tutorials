// package main

// import (
// 	"fmt"
// )

// // type rectangle struct {
// // 	height, width float64
// // }

// // type circle struct {
// // 	radius float64
// // }

// // func (c circle) area() float64 {
// // 	return math.Pi * math.Pow(c.radius, 2)
// // }

// // func (c circle) perimeter() float64 {
// // 	return math.Pi * c.radius * 2
// // }

// // func (c circle) volume() float64 {
// // 	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
// // }

// // func (r rectangle) area() float64 {
// // 	return r.height * r.width
// // }

// // func (r rectangle) perimeter() float64 {
// // 	return 2 * (r.width + r.height)
// // }

// // func (r rectangle) diagonalLength() float64 {
// // 	return math.Sqrt((math.Pow(r.height, 2)) + (math.Pow(r.width, 2)))
// // }

// // func (c circle) getColor() string {
// // 	return "green"
// // }

// // type shape interface {
// // 	area() float64
// // 	perimeter() float64
// // }

// // type object interface {
// // 	volume() float64
// // }

// // //
// // //
// // //

// // //** EMBEDDEDED INTERFACE **//
// // // There is no inheritance in golang, but there is room for emebedding a interface into another.
// // // The below interface geometry embeds the shape and object interfaces, so it will have access to their
// // // methods.

// // type geometry interface {
// // 	shape
// // 	object
// // 	getColor() string
// // }

// // func calc(s shape) {
// // 	fmt.Printf("Shape is %v\n", s)
// // 	fmt.Printf("Area of the %s is %.2f\n", s, s.area())
// // 	fmt.Printf("Perimeter of the %s is %.2f\n", s, s.perimeter())
// // }

// // //
// // //
// // //

// // //** EMPTY INTERFACE **//
// // // This is an interface that has no methods.
// // // it can hold any values of any type

// // type emptyInterface interface{}

// // func main() {
// // 	//An interface is collection of method signatures that an object (most times a named type) can implement
// // 	circle1 := circle{radius: 5}
// // 	rec1 := rectangle{height: 2, width: 10}

// // 	fmt.Printf("The area and perimeter of the cirlce with the radius 5 is: %.2f and %.2f\n", circle1.area(), circle1.perimeter())

// // 	// interface is to allow us to no repeat ourselves, by implementing interface, we can name all methods that a variable can implement
// // 	// without writing separate functions for each and every types

// // 	// USING INTERFACE
// // 	calc(circle1)
// // 	calc(rec1)

// // 	//
// // 	//
// // 	//

// // 	//** INTERFACES TYPES AND VALUES **//
// // 	// Interfaces has dynamic types that changes in runtime.

// // 	var s shape
// // 	var box shape
// // 	fmt.Printf("%T\n", s) // <nil>

// // 	ball := circle{radius: 3.5}
// // 	// ball as a variable has the ability to perform the two functions in the shape type
// // 	// which makes the s varaibale be able to be asigned the variable ball

// // 	s = ball
// // 	fmt.Printf("%T\n", s) // main.circle // its type has changed to circle

// // 	room := rectangle{height: 5, width: 10}
// // 	// same here too
// // 	s = room
// // 	fmt.Printf("%T\n", s) //main.rectangle // its type has changed to rectangle

// // 	//
// // 	//
// // 	//

// // 	//**TYPE ASSERTIONS AND TYPE SWITCHES**//

// // 	// Type assertion happens when a variable of a type interface is trying to access
// // 	// another method not suitbale for it or not in the interface

// // 	s = circle{radius: 2.5}
// // 	box = rectangle{height: 3, width: 10}

// // 	// writing the below expression will cause error.
// // 	//s.volume() // s being a circle type variable only has 2 methods to perform from its interface.

// // 	// The only way s can access the volume method is by using type assertion.

// // 	//fmt.Println(s.(circle).volume())

// // 	ballVolume, assertion_ok1 := s.(circle)
// // 	boxDia, assertion_ok2 := box.(rectangle)

// // 	if assertion_ok1 {
// // 		fmt.Printf("The volume of the ball is %.2fv³\n", ballVolume.volume())
// // 	}
// // 	if assertion_ok2 {
// // 		fmt.Printf("The diagonal length of the box is %.2f\n", boxDia.diagonalLength())
// // 	}

// // 	//
// // 	//
// // 	//

// // 	// TYPE SWITCHES

// // 	s = rectangle{height: 4, width: 16}
// // 	switch value := s.(type) {
// // 	case circle:
// // 		fmt.Printf("%#v has circle type\n", value)
// // 	case rectangle:
// // 		fmt.Printf("%#v has rectangle type\n", value)
// // 	}

// // 	//
// // 	//
// // 	// EMBEDDED INTERFACE EXAMPLE //

// // 	var geoSample geometry
// // 	geoS := circle{radius: 3.4}
// // 	geoSample = geoS

// // 	fmt.Println(strings.Repeat("-", 20))

// // 	fmt.Println(geoSample.getColor())
// // 	fmt.Println(geoSample.area())
// // 	fmt.Println(geoSample.volume())
// // 	fmt.Println(geoSample.perimeter())

// // 	//
// // 	//
// // 	//

// // 	//** EMPTY INTERFACE **//

// // 	var empty interface{}

// // 	empty = 5
// // 	fmt.Println(empty) // 5

// // 	empty = "Golang"
// // 	fmt.Println(empty) // "Golang"

// // 	empty = []int{1, 2, 5}
// // 	fmt.Println(empty) // l

// // 	// using type assertions, we can get the dyamic value of the interface

// // 	fmt.Println(len(empty.([]int))) // 3

// // 	fmt.Println(strings.Repeat("-", 20))

// // 	//EMpty interface is useful for running program that contains unknown types.

// // 	type person struct {
// // 		id       int
// // 		userInfo interface{}
// // 	}

// // 	firstPerson := person{id: 1, userInfo: "His name"}
// // 	secondPerson := person{id: 2, userInfo: 50}
// // 	thirdPerson := person{id: 3, userInfo: []int{4, 5, 9}}
// // 	fmt.Println(firstPerson, secondPerson, thirdPerson)

// // 	// Here the userInfo interfcae was able to hold any data given to it, which makes empty interfcae powerful.

// type money float64

// type book struct {
// 	price float64
// 	title string
// }

// func (m money) print() {
// 	fmt.Printf("%.2f\n", m)
// }

// func (m money) printStr() string {
// 	return fmt.Sprintf("%.2f", m)
// }

// func (b book) vat() float64 {
// 	return 0.09 * b.price
// }

// func (b *book) discount() float64 {
// 	b.price = b.price - (b.price * 0.1)
// 	return b.price
// }

// //
// //
// // INETERFACES QUESTIONS //

// type vehicle interface {
// 	License() string
// 	Name() string
// }

// type car struct {
// 	licenceNo string
// 	brand     string
// }

// func (c car) License() string {
// 	return c.licenceNo
// }

// func (c car) Name() string {
// 	return c.brand
// }

// func main() {
// 	var m1 money = 5.6
// 	book1 := book{price: 300.67, title: "A new book"}

// 	m1.print()
// 	fmt.Println(m1.printStr())

// 	fmt.Println(book1.vat())
// 	fmt.Println(book1.discount())

// 	//
// 	//
// 	//

// 	// ** INTEFACES QUESTIONS **//

// 	var vehInt vehicle
// 	car1 := car{licenceNo: "AG786BHD", brand: "AUDI"}

// 	vehInt = car1

// 	fmt.Println(vehInt.License())
// 	fmt.Println(vehInt.Name())

// }

package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{} = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
