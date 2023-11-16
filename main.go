package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Starting Program")
	distance := 60.8
	fmt.Printf("The distance in miles is %f \n", distance*0.621)
	fmt.Println("The distance in miles is", distance*0.621)
	fmt.Println("The time in seconds is", secondsInHour)
	fmt.Println("Program ran successfully!!!")
}
