package main

import "fmt"

func main() {

	const (
		North = iota // Iota is a data type that increases by 1
		South
		West
		East
	)

	const (
		a = iota * 2
		b
		c
		d
	)

	const (
		jun = (iota) + 6
		jul
		aug
	)

	// const (
	// 	x = (iota) - 2
	// 	y = (iota) - 5
	// 	z = (iota) - 7
	// )

	//another way

	// const (
	// 	x = -(iota + 2)
	// 	_
	// 	y
	// 	z
	// )

	// fmt.Println(a, b, c, d) // 0 , 2,  4, 6

	// fmt.Println(x, y, z) //

	const x = 7
	const y float64 = 3.1
	fmt.Println(x * y) // There wont be error because
	//x is an untyped constant and it will get its type from the first expression where it's used.
}
