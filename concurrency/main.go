package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1() {
	fmt.Println("f1 goroutine started")
	for i := 0; i < 5; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 goroutine ended")
}

func f2() {
	fmt.Println("f2 goroutine started")
	for i := 5; i < 9; i++ {
		fmt.Println("f2 i=", i)
	}
	fmt.Println("f2 goroutine ended")
}

func main() {
	// CONCURRENCY.
	// Spawning Goroutines

	// Printing out some system information

	// fmt.Println("No of Cores", runtime.NumCPU())
	// fmt.Println("No of Goroutines", runtime.NumGoroutine())
	// fmt.Println("OS", runtime.GOOS)
	// fmt.Println("Architecture", runtime.GOARCH)
	// fmt.Println("No of OS threads in parallel", runtime.GOMAXPROCS(0))

	// A gotoutine is a function called with go keyword before it name.

	f1()
	go f2()
	fmt.Println("No of Goroutines", runtime.NumGoroutine())

	time.Sleep(time.Millisecond) // But this method isn't logical in real world application.

	//
	//
	//

	// **WAIT GROUPS**//

	// STEP 1: instantiate sync.WaitGroup to an instance
	wg := sync.WaitGroup{}
	// OR//
	var waitG sync.WaitGroup

	//STEP 2: call the add of n method before attempting to execute the goroutine
	wg.Add(1) // to specify the number of goroutines to wait for

	// STEP 3: change the function definition
}
