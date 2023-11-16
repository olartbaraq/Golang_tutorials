// ARRAYS

package main

import "fmt"

func main() {
	// nums := [...]int{30, -1, -6, 90, -6}
	// count := 0

	// for _, num := range nums {
	// 	if num%2 == 0 && num > 0 {
	// 		count++
	// 	}
	// }
	// fmt.Printf("The number of positive integers is %d", count)

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[1] = 10.10

	fmt.Println(myArray)
}

// func array() {
// 	// Introduction

// 	// var numbers [3]float64 // here we have an array of float64 numbers with 3 elements intialized to zero

// 	// fmt.Printf("%v\n", numbers)

// 	// Declaration of array

// 	// var age = [5]int{10, 12, 35, 48, 26}
// 	// _ = age

// 	// var salary = [2]float64{255.56, 300.76}
// 	// _ = salary

// 	// names := [2]string{"Olatunde", "Mubaraq"}
// 	// _ = names

// 	// ellipses way
// 	// age := [...]int{10, 12, 3, 6, 26} // using thre ... inside the bracket

// 	// fmt.Printf("the number of elements in the array is %d and the first element is %d \n", len(age), age[0])

// 	// Declaring on multiple lines
// 	// You must end the last element of the array with a comma

// 	//ARRAY OPERATIONS

// 	// modifying the array elements

// 	// numbers := [4]int{5, 6, 7, 8}
// 	// numbers[0] = 10 // the first element, index 0 of the array has now being chnaged to 10 from 5.

// 	// accessing and iterating the array elements

// 	// 1- Using the range key word

// 	// for index, number := range numbers {
// 	// 	fmt.Printf("the number %d is at index %d\n", number, index)
// 	// }

// 	//2- using the for loop on the length of the array

// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Println("index", i, "number", numbers[i])
// 	// }

// 	// WORKING WITH MULTIDIMENSIONAL ARRAY

// 	figures := [3][4]int{
// 		{1, 2, 3, 4},
// 		{4, 5, 6, 7},
// 		{7, 8, 9, 10},
// 	}
// 	//fmt.Println(figures)

// 	// ARRAY EQUALITY

// 	// two arrays are equal if they have the same length, elements and in the same order

// 	numbers := figures

// 	fmt.Println(figures == numbers) // prints true

// 	//ARRAYS WITH KEYED ELEMENTS

// 	grades := [3]int{
// 		0: 6,
// 		2: 3,
// 		1: 6,
// 	}
// 	fmt.Println(grades)

// 	accounts := [3]int{2: 50} // it means the element at index 2 is 50, while the element at index 0 and 1 is 0.
// 	_ = accounts

// 	games := [...]string{
// 		5: "God of War",
// 	}

// 	cities := [...]string{
// 		5:        "Lagos",
// 		"Ibadan", // this is going to be the last element in the array.
// 		4:        "Riyaad",
// 	}

// 	fmt.Println(cities, cities[0], len(cities))

// 	fmt.Println(len(games), games) // prints an array of strings with its element at index 5 is "God of War" while
// 	// the element at index 0 to 4 is empty string "".

// 	// if you have an array , you can specify the index of some element in the array and leave
// 	// out others to take the default value

// 	days := [7]bool{4: true, 1: true} // here elements in indeces 1 and 4 will be true while others are false.
// 	fmt.Println(days)
// }
