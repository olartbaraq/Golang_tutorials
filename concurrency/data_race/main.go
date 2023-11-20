package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// n := 0
	// const gr = 100

	// wg := sync.WaitGroup{}

	// wg.Add(2 * gr)

	// for i := 0; i < gr; i++ {
	// 	go func() {
	// 		time.Sleep(time.Second / 10)
	// 		n++
	// 		wg.Done()
	// 	}()

	// 	go func() {
	// 		time.Sleep(time.Second / 10)
	// 		n--
	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()

	// fmt.Println(n)

	// FOr the goroutine, the desired value for n is zero but die to data sharing because we wont know the
	// goroutine that will finish execution first, so we wont know the final value of n. n will be incremeitng or decrementing at times.
	// The value of n changes depending on the value of the goroutine that finishes execution first.
	// Thats why it's called data race

	//
	//
	//

	//** GO RACE DETECTOR **//

	// to see where the issue is run

	//go run -race [filename].go

	//
	//
	//

	//** MUTEXES **//

	//Mutex is used to solve the problem of data races.

	n := 0
	const gr = 100

	wg := sync.WaitGroup{}

	wg.Add(2 * gr)

	//STEP:1 instantiate a Mutex value
	m := sync.Mutex{}

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)

			//STEP:2 Block access to the shared value
			m.Lock()
			n++
			//STEP:3 -  Unlock the variable after it's incremented
			m.Unlock()

			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			//m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(n)
}
