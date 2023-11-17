package main

import (
	"fmt"
	"strings"
)

func accPtr(ptr *int, val int) *int {
	*ptr = val
	return ptr
}

type Product struct {
	name  string
	price float64
}

func changeStruct(p Product) {
	// lets modify the struct
	p.name = "Hp Envy pro"
	p.price = 1999.99
}

func changeStructByPointer(p *Product) {
	// lets modify the struct
	p.name = "Hp Envy pro"
	(*p).price = 1999.99 // still equivalent to using p.price = 1999.99
}

func updateSlice(s []int, a int) {
	for i := range s {
		s[i] = a
	}
}

func updateMap(m map[string]int, y int) {
	for k, v := range m {
		m[k] = y
		m["new"] = v
	}
}

func main() {

	// Variable is just a convenienent name used to be label memory location in CPU.

	// POINTER
	// A pointer is a variable that stores the memory address of another variable.
	//Say

	var num int = 5
	// num is the label (name) used to store the memory location of the value 5 of type int.

	numAdd := &num
	fmt.Println(numAdd)

	//numAdd is the address of num or nil is num hasn't been initialized
	//numAdd stores the address of num in memory
	// i.e numAdd is a pointer to num.

	//NB: Pointers have the power to change the data they are pointing to.
	*numAdd = 7 // here we have updated the value of num to 7 without knowing the variable name numAdd pointed to
	fmt.Println(num, numAdd, &numAdd)

	// Unlike C; Golang has no pointer arithmetric

	// Declaring a pointer to a type which hasn't been initialized
	var ptr *float64
	// OR //
	fmt.Println(ptr)
	//

	// using the new built -in function take a type as argument and allocates enough memory
	// to accomodate the value of that type and returns a pointer to it.
	// This will have addresss and not nil.

	var p = new(int)
	j := new(float64)
	x := 100
	y := 100.6
	p, j = &x, &y
	fmt.Println(p, j)

	fmt.Println(strings.Repeat(("-"), 20))

	//
	//
	//
	// **DEREFERENCING OPERRATOR**//
	*p = 90 // Here am derefencing (trying to get) the value pointer p points to
	//i.e if p points to the address , then *p holds the value stored in the address
	// which means if *p = 90 then x= 90
	fmt.Printf("p which is the address is %p\n *p = %v and the value of x is %v\n", p, *p, x)

	//
	//
	// **POINTER TO POINTER**//

	pintr := new(int) // pintr is a pointer to type Int
	pInt := &pintr    // pIntr is a pointer to pointer pintr which is a pointer to type Int

	fmt.Printf("pintr is %p and pInt is %p and value of type Int is %v and the value of type Int is %v\n", pintr, pInt, *pintr, **pInt) // <Ox...>, <0x...>, 0, 0
	intV := 300
	pintr = &intV
	fmt.Printf("pintr is %p and pInt is %p and value of intV is %v\n", pintr, pInt, *pintr) // <Ox...>, <0x...>, 300

	**pInt = 400                                    // Here i have used pInt which is the pointer to pointer to type int to change the value of intV to 400
	fmt.Printf("New value of intV is %v\n", **pInt) // 400

	fmt.Println(strings.Repeat(("-"), 20))

	//
	//
	// **COMPARING POINTERS.**//

	// The zero value of any pointer to any type is nil.
	// pointers comapred to each other if and only if they point to the same varaible or are nil

	var cc = new(int) // cc will have address
	var dd *int       // dd will have no address and be nil
	//"cc and dd are pointer to type int"
	fmt.Println(cc == dd) // false

	var ee *float64
	var ff *float64
	fmt.Println(ee == ff) // true because both ee and ff are nil

	numX := 560.67

	ee = &numX
	ff = &numX

	fmt.Println(ee == ff) // true because both ee and ff points to the same variable

	numY := 560.67
	ff = &numY
	fmt.Println(numX == numY) // true because they point to the same **VALUE**
	fmt.Println(ee == ff)     // false because they point to different variables even if those variables were equal.

	fmt.Println(strings.Repeat(("-"), 20))

	//
	//
	//
	// **PASSING POINTERS TO FUNCTIONS AND RETURNING POINTERS FROM FUNCTIONS**//
	z := 8
	pp := &z
	fmt.Println(pp, z, *pp) // <0x...>, 8, 8
	accPtr(pp, 10)
	fmt.Println(pp, z, *pp) //<Ox..>, 10, 10

	fmt.Println(strings.Repeat(("-"), 20))

	//
	//
	//
	// **PASSING OR STRUCT TO FUNCTIONS**//

	product1 := Product{name: "Dell latitude 7490", price: 999.99}
	changeStruct(product1)
	fmt.Printf("Nothing changes to product 1 without using pointer %+v\n", product1) // Here nothing changes in the struct
	changeStructByPointer(&product1)
	fmt.Printf("Product 1 has now been updated to %+v after passing it by pointer\n", product1) // Here nothing changes in the struct

	// NB: to update the value of int, string, float , struct by functions, we need to pass the pointer to them as values and not their actual values.
	// but its different from Maps and slice.

	fmt.Println(strings.Repeat(("-"), 20))

	//
	//
	//
	// **PASSING SLICES OR MAPS TO FUNCTIONS**//

	//When a function chnages a map or slice, the actual values changes, it doesn't need to be a pointer to
	// the map or slice

	slices1 := []int{2, 4}
	fmt.Printf("%+v\n", slices1)
	updateSlice(slices1, 10) // This function will update the two value slices to 10,10
	fmt.Printf("%#+v\n", slices1)

	fmt.Println(strings.Repeat(("-"), 20))
	// Updating a map
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("%+v\n", map1)
	updateMap(map1, 5)
	fmt.Printf("%+v\n", map1) // This function will update the two value slices to ["new": 2, "two" : 5, "one": 5]

}
