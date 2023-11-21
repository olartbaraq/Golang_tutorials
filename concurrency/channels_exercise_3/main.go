package main

import (
	"fmt"
	"math"
)

// func power(x int, c chan float64) {
// 	sqX := math.Pow(float64(x), 2)
// 	c <- sqX
// }

func main() {
	// ch := make(chan string)

	// go func(str string, c chan string) {
	// 	c <- str
	// }("mubby", ch)

	// fmt.Println(<-ch)

	ch := make(chan float64)

	// for i := 1; i <= 50; i++ {
	// 	go power(i, ch)
	// 	re := <-ch
	// 	fmt.Printf("power of %d is %.2f\n", i, re)
	// }

	// Using function literal

	for i := 1; i <= 50; i++ {
		go func(x int) {
			sqX := math.Pow(float64(x), 2)
			ch <- sqX
		}(i)
		re := <-ch
		fmt.Printf("power of %d is %.2f\n", i, re)
	}
}
