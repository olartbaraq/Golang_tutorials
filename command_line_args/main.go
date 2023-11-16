package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args:", os.Args)
	fmt.Println("First Argument:", os.Args[1])
	fmt.Println("Second Argument:", os.Args[2])
	fmt.Println("Third Argument:", os.Args[3])
	fmt.Println("Total Argument:", len(os.Args))

	// For Loops in golang

	// for i:= 0; i < len(os.Args); i++ {
	// 	result, err := strconv.Atoi(os.Args[i]); err == nil {

	// 	}
	// 	fmt.Printf("%s Argument: ", os.Args[i])
	// }
}
