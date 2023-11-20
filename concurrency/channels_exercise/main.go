package main

import (
	"fmt"
	"strings"
)

// func f1(n int, ch chan int) {
// 	// send the value of n to the channel
// 	ch <- n
// }

func factorial(n int, ch chan int) {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func main() {

	// c := make(chan int)
	// // create a new channel for bi- directional operations

	// go f1(10, c)
	// // run a goroutine of function f1 with arguments

	// new := <-c
	// // receive from the channel to a new value

	// fmt.Println(new)

	factCh := make(chan int)
	defer close(factCh)

	go factorial(5, factCh)
	f := <-factCh
	fmt.Println(f)

	fmt.Println(strings.Repeat("-", 20))

	for i := 1; i <= 20; i++ {
		go factorial(i, factCh)
		eachF := <-factCh
		fmt.Println(eachF)
	}

	fmt.Println(strings.Repeat("-", 20))

	// Using function literal or anonymous functions

	for i := 1; i <= 20; i++ {
		go func(n int, ch chan int) {
			result := 1
			for i := 2; i <= n; i++ {
				result *= i
			}
			ch <- result
		}(i, factCh)

		eachF := <-factCh
		fmt.Println(eachF)
	}

}
