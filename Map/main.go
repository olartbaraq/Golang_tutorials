package main

import (
	"fmt"
)

func main() {
	// MAP - a collection type like an array or slice but it saves in key/value pairs format.
	// just like dictionary in python and object in javascript

	// It's an unordered data structure

	// add, get , delete are operations done at constant time

	// For a Map to be valid in Go, the Keys must be unique and be of the same type
	// and the values must of the same type but may not be unique.

	// We can only compare a map to nil just like a slice.

	//NB: its recommended not to use float64 as key in Map data structure

	//
	//

	// CREATING A MAP

	//var employees map[string]string // employees is a map that contains the type string as the key and values being
	// initalized to nil. i.e the map has not elements.
	//fmt.Println(employees)
	//_ = employees

	// To actually populate a map, we can use the make function or standard way.

	//1- STANDARD WAY:
	// people := map[string]float64{} // this is a map that is empty but initialized, it's not nil.
	// people["John"] = 21.4
	// people["Akeem"] = 10.3

	// fmt.Println(people)

	//2 - MAKE FUNCTION METHOD
	// names := make(map[string]int)
	// names["John"] = 3
	// names["Akeem"] = 5
	// fmt.Println(names)

	//3 - POPULATING WITHIN THE INITIALIZATION
	// currency := map[string]float64{
	// 	"NGR": 100.34,
	// 	"USD": 200.56,
	// 	"EUR": 100,
	// 	"FRF": 90.4,
	// }

	// currency["EUR"], currency["CAF"] = 100.50, 90

	// fmt.Println(currency)

	//NB: trying to look up in the map a key that doesn't exist will return the default value of the value type(float64).
	//fmt.Println(currency["BNG"]) // 0

	// But to actually find out if a key exists or not, we have to use a (comma oka idiom). This iwll check for the
	// corresponding value of the key and also return a bool if the key exists or not
	// value, isExisting := currency["BNG"]
	// fmt.Println(value, isExisting) // 0, false.
	// SO seeing false means the key doesn't exist in the map.

	//
	//
	//

	//LOOPING OVER A  MAP
	// for k, v := range currency {
	// 	fmt.Printf("The key is %q and value is %v \n", k, v)
	// }

	//
	//
	//

	//DELETING A KEY FROM A MAP.
	// deleting a key not in the map will throw no error and nothing will be deleted
	// delete(currency, "BNG") // no erorr and nothing is deleted
	// delete(currency, "CAF") // CAF key is deleted
	// fmt.Println(currency)

	//
	//
	//

	//COMPARRING MAPS.
	//Maps can only be compared to nil, i.e we cant compare maps.
	// The only somewhat way to compare maps is to compare their string representation
	// using the Sprintf(), but the catch is the key/value pair must be of type string.

	// a := map[string]string{
	// 	"NGR": "100.34",
	// 	"USD": "200.56",
	// 	"EUR": "100",
	// 	"FRF": "90.5",
	// }
	// b := map[string]string{
	// 	"NGR": "100.34",
	// 	"USD": "200.56",
	// 	"EUR": "100",
	// 	"FRF": "90.5",
	// }

	// s1 := fmt.Sprintf("%s", a)
	// s2 := fmt.Sprintf("%s", b)
	// fmt.Println(s1 == s2) //  true

	//
	//
	//

	//MAP HEADER AND CLONING MAPS
	//Copying a map to another varaible(of type map) shares the same map header. the two maps will point to the same
	// map header in memory just like an array/slice when created using the slicing function.

	// friends := map[string]int{
	// 	"MB":   14,
	// 	"Tola": 20,
	// }

	// oldF := friends

	// friends["MB"] = 56

	// fmt.Println(oldF)    // ["MB":56, "Tola":20]
	// fmt.Println(friends) // ["MB":56, "Tola":20]

	// oldF is not a copy of friends, instead it references the friends map structure in memory.

	//To actaully copy a map to another map, we need to initialize the new map and then use for loop to copy each
	//element into the new map.

	coun_Rank := map[string]int{
		"Nigeria": 1,
		"Angola":  2,
		"Romania": 3,
	}

	newFr := make(map[string]int)

	for k, v := range coun_Rank {
		newFr[k] = v
	}

	fmt.Println(newFr) // this will have the key/value pairs of friends which doesn't point to the same map headers as friends.
	coun_Rank["Romania"] = 6

	fmt.Println(newFr)
	fmt.Println(coun_Rank)

	// Updating the value of the key Romania doesn't affect the copied map newFr.

}
