package main

import "fmt"

func main() {
	// Copying a slice to another slice

	src := []int{1, 2, 3, 4}

	dst := make([]int, len(src)-1)

	final := copy(dst, src)
	_ = final
	fmt.Println(dst) // [1,2,3]
	// dst- first 3 elements will be copied into dst from src

	//SLICE EXPRESSIONS

	// slicing an array becomes a slice , same as slice
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:4] // b = []int{1, 2, 3, 4} i.i it will start from 0 and end in before stop
	// in other words b := array[start:stop] but the stop will be excluded
	c := a[1:3]
	fmt.Println(b, c)

	d := []int{1, 2, 3, 4, 5, 6, 7}
	e := d[1:4] // e = [2, 3, 4]
	fmt.Println(e)

	//Also when using the start and stop parameters, a missing start defaults to 0 and a missing stop defaults to the last
	// element of the array/slice
	//NB: with a missing stop parameter, the new slice will contain the last element of the old array/slice

	f := d[:4] // [1,2,3,4]
	g := d[1:] // [2,3,4,5,6,7]
	h := d[:]  // [1,2,3,4,5,6,7] same as d[0:len(d)]
	fmt.Println(f, g, h)

	//APPENDING NEW ELEMENTS TO A SLICE

	// f = append(f[:2], 100)
	// fmt.Println(f) // [1,2,100]

	g[2] = 100
	g = append(g, 100)
	fmt.Println(g, d, e, f, h)

	//BACKING ARRAY

	//NB. When a new slice is created by slicing another array/slice e.g e, f, g, h.
	//Whenever an element is appended or updated in the new slice, it will override that existing element of all other new slices created
	// and also the original slice.
	// upon updating index 2 to 100 at slice g and appending 100 to slice g, all other slice that stern from slice (a) and slice (a) inclusive
	//will have their index 3 updated to 100 respectively.

}
