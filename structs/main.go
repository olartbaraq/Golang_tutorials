package main

import (
	"fmt"
	"strings"
)

func main() {
	// // struct is a schema, defined data structure. it's fixed at compile time.

	// //
	// //
	// // CREATING A STRUCT.

	// type book struct {
	// 	title  string
	// 	author string
	// 	year   int
	// }
	// // book is a type struct, and title, author and year are fields of the struct with type string,
	// //string and int respectively.

	// type car struct {
	// 	name, model string
	// 	year        int
	// 	price       float64
	// }

	// // book1 := book{"title1", "author1", 1920}
	// car1 := car{"toyota", "camry", 2023, 4500}
	// _ = car1

	// // book1, car1 are structs while book, car are struct type

	// // But the above way can be prone to error if the fields were to be placed wrongly
	// //instead we tie the field name to the appropraite value

	// book2 := book{author: "author3", title: "title3", year: 2003}

	// //fmt.Println(book1, book2, car1)

	// //NB: if some fields defined in the struct were to be ommitted when initializing
	// //to a variable, the default value of the data type will be placed automatically
	// // and not an error, which can then now be updated.

	// //
	// //
	// //

	// // RETRIEVING AND UPDATING STRUCT FIELDS
	// // To retrieve and update the value of a struct field, we use the dot notation.

	// // RETRIEVING A FIELD VALUE
	// fmt.Println(book2.author) // author3

	// //UPDATING A FIELD VALUE
	// book2.author = "new Author"

	// fmt.Println(book2) // {author: "new Author", title: "title3, year: 2003}

	// //
	// //
	// //

	// // COMPARING STRUCT FIELDS
	// //  Two structs are equal if their corresponding fields are equal

	// book4 := book{author: "new Author", title: "title3", year: 2003}
	// fmt.Println(book2 == book4) // true . All corresponding fields are equal

	// //
	// //
	// //

	// //COPYING A STRUCT TO ANOTHER
	// // Unlike array, slice where = just create a new slice that points to the same header in memory,
	// // in struct, a brand new struct is created and modifying one doesn't update the other since they are
	// // different in memory.

	// newBook := book4
	// newBook.author = "Nadine"

	// fmt.Println(newBook, book4) // no field chnages, just the author of the new book.

	// //
	// //

	// // ANONYMOUS STRUCTS AND FIELDS

	// // 1- making a new struct with without specifying the type before hand

	// diana := struct {
	// 	firstName, lastName string
	// 	age                 int
	// }{
	// 	firstName: "Dinah",
	// 	lastName:  "Madini",
	// 	age:       23,
	// }
	// fmt.Printf("%+v\n", diana)

	// fmt.Println(strings.Repeat("#", 10))

	// //2 - creating a new struct without a type will make the fields type be the name of the
	// //fields.

	// type name struct {
	// 	string
	// 	float64
	// 	int
	// }

	// person1 := name{"Nadine", 13.56, 25}
	// fmt.Printf("%+v\n", person1)

	// fmt.Println(strings.Repeat("#", 10))

	// //3 - creating a new struct with and without a type
	// type laptop struct {
	// 	name string
	// 	float64
	// 	model string
	// }

	// laptop1 := laptop{name: "MacBook Pro M1", model: "Pro M1", float64: 1999.99}
	// fmt.Printf("%+v\n", laptop1)

	// //
	// //
	// //

	// EMBEDDED STRUCTS

	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact // Embedded struct
	}

	employee1 := Employee{
		name:   "Nadine",
		salary: 180,
		contactInfo: Contact{
			email:   "nadine@gmail.com",
			address: "No 7, Nadina St, New York, Ibadan",
			phone:   1111111111,
		},
	}

	fmt.Printf("%+v\n", employee1)
	fmt.Println(strings.Repeat("#", 20))

	//Accessing some fields

	fmt.Printf("%+v\n", employee1.contactInfo.email)
	fmt.Println(strings.Repeat("#", 20)) // "nadine@gmail.com"

}
