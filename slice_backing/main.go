package main

import "fmt"

func main() {
	// Since we cannot create a new slice from an existing slice/array, because whatever elements changed will
	// update in other slices, the only way to do so is to use the append function.

	// var cars = []string{"Ford", "BMW", "Toyota", "Audi"}
	// var newCars = []string{}
	// newCars = append(newCars, cars...)
	// newCars2 := cars[0:3]

	// newCars2[0] = "Mercedes"

	// cars[0] = "Honda"
	// newCars[0] = "Mazda"

	//fmt.Println(newCars, cars, newCars2)

	// since newCars wasn't created using the slicing method like
	//newCars =: cars[0:3]. updating any elements in cars and newCars will only effect the updated slice.

	// so here, updating newCars2 will update the first element to of cars to also "Mercedes" because newcars2 was
	//created using the slicing method but it wont affect newCars because it was created using the append funtion.

	// APPEND, LENGTH AND CAPACITY IN SLICES
	//Append function is not always very effective when creating a new slice.

	// var nums []int // a nil slice
	// var newArray []int
	// fmt.Printf("Length: %d, Capacity: %v, Header : %p\n", len(nums), cap(nums), &nums)

	// nums = append(nums, 1, 2, 3)
	// newArray = append(newArray, nums...)
	// fmt.Printf("Length: %d, Capacity: %v, Header : %p\n", len(nums), cap(nums), &nums)
	// fmt.Printf("Length: %d, Capacity: %v, Header : %p\n", len(newArray), cap(newArray), &newArray)

	//EXERCISE

	mySlice := []float64{1.2, 5.6}

	//mySlice[2] = 6

	a := 10
	mySlice[0] = float64(a)

	//mySlice[1] = 10.10

	mySlice = append(mySlice, 10.10, float64(a))

	fmt.Println(mySlice) // [10, 5.6,10.10, 10]

	//

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	var newYears []int

	newYears = append(newYears, years[0:3]...)
	newYears = append(newYears, years[len(years)-3:]...)

	fmt.Println(newYears)
}
