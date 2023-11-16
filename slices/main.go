package main

import "fmt"

// func slice1() {
// 	//SLIES

// 	// its just like an array but has a dynamic size i.e it can shrink or grow at runtime

// 	//1- Declaring Slices and basic operations

// 	// cities := [...]string{} // array
// 	// var countires []string  // slice

// 	// fmt.Printf("cities type is %T\n", cities)
// 	// fmt.Printf("cities type is %T\n", countires)

// 	// numbers := []int{2, 4, 1, 7, 8}
// 	// _ = numbers

// 	//2- Declaring Slices using the make function, it takes 2 args (type of slice and the num of elements)

// 	// nums := make([]int, 2) // slice with 2 numbers initialised to zeros
// 	// _ = nums

// 	//3 Using the type data structure

// 	// type names []string

// 	// friends := names{"Mubaraq", "Olatunde"}

// 	// fmt.Println(friends)

// 	//HOW TO ITERATE OVER A SLICE
// 	// for i, value := range friends {
// 	// 	fmt.Println("index:", i, " value", value)
// 	// }

// 	// age := [...]int{1, 2, 3, 4}
// 	// figure := []int{1, 2, 3, 4}

// 	// fmt.Println(figure == age) // mismatched type of [4]int (array) to []int (slice)

// 	//COMPARING SLICES

// 	var n []int
// 	fmt.Println(n == nil) // true

// 	var m = []int{}
// 	fmt.Println(m == nil) // false beacuse its not empty

// 	// Slices can only be compaired to nil, we can't compare 2 slices even if they have the same
// 	// length, elememts and are in same order

// 	a, b := []int{1, 2, 3, 4}, []int{1, 2, 3, 4}
// 	// fmt.Println(a == b) // invalid ops: slice can only be compaired to nil

// 	//Slice can only be compared to each other by iterating over the slices and compare elements by elements

// 	// var count = 0
// 	// for _, valA := range a {
// 	// 	for _, valB := range b {
// 	// 		if valA == valB {
// 	// 			count++
// 	// 		}
// 	// 	}
// 	// }
// 	// if count == len(a) && count == len(b) {
// 	// 	fmt.Println("Both slices are equal", count)
// 	// } else {
// 	// 	fmt.Println("Both slices are not equal", count)
// 	// } // THis method is O(n squared)

// 	// METHOD 2
// 	eq := true

// 	for i, valA := range a {
// 		if valA != b[i] {
// 			eq = false
// 			break
// 		}
// 	}

// 	if !eq {
// 		fmt.Println("Both slices are not equal")
// 	} else {

// 		fmt.Println("Both slices are equal")
// 	} //THis method is O(n squared)

// }

func main() {
	// Appending a Slice and Copying Slices

	var numbers = []int{2, 3}

	//1- adding a new element to slice at the end position using the append function
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	//2 - appending a slice to another slice
	n := []int{200, 300}
	numbers = append(numbers, n...) // put ... at the end of slice you want to append
	fmt.Println(numbers)

	//3 - Copying a slice to another slice
	ss := []int{20, 30}
	dd := make([]int, len(ss))
	ee := copy(dd, ss)
	fmt.Println(ee)
}
