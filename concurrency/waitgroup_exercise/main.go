package main

import (
	"fmt"
	"math"
	"sync"
)

// func sayHello(n string, wg *sync.WaitGroup) {
// 	fmt.Printf("Hello, %s!\n", n)
// 	wg.Done()
// }

// func sum(num1 float64, num2 float64, wg *sync.WaitGroup) {
// 	sum := num1 + num2
// 	fmt.Printf("%.2f\n", sum)
// 	wg.Done()
// }

func main() {
	wg := sync.WaitGroup{}

	wg.Add(50)

	for i := 100; i <= 149; i++ {
		go func(x float64, wg *sync.WaitGroup) {
			y := math.Sqrt(x)
			fmt.Printf("sqaure root of %f is %f\n", x, y)
			wg.Done()
		}(float64(i), &wg)

	}

	// go sum(5.5, 5.4, &wg)
	// go sum(5, 2.8, &wg)
	// go sum(6.5, 9.4, &wg)

	//go sayHello("Mr. Wick", &wg)

	wg.Wait()

	fmt.Println("all tasks done")
}
