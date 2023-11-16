package main

import "fmt"

type duration int

func main() {
	type salary float64       // a new defined data type with its underlying type to be float64
	type oldSalary salary     // a new defined data type with its underlying type to be float64 and not salary
	type veryOldSalary salary // a new defined data type with its underlying type to be float64 and not salary

	x := salary(9.45)        // x is a varaible of type salary but its underlying type float64
	y := oldSalary(9.45)     // y is a varaible of type oldSalary but its underlying type float64
	z := veryOldSalary(9.45) // z is a varaible of type veryOldSalary but its underlying type float64

	fmt.Printf("salary is of type %T \n", x)
	fmt.Printf("salary is of type %T \n", y)
	fmt.Printf("salary is of type %T and its value is %.2f \n", z, z)

	//TYPE ALIASES

	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))

}
