package main

//import "fmt"

func main() {
	var x uint8 = 255
	//x++ // overflows occur and x = 0.
	_ = x

	// but if expression leads to overflow, it shows error in compile time

	//y := float32(5)
	//var y float32 = 5
}
