package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(200)

	m := sync.Mutex{}

	balance := 100

	for i := 0; i < 100; i++ {

		go func(b *int, n int, wg *sync.WaitGroup) {
			m.Lock()
			defer m.Unlock()
			*b += n
			wg.Done()
		}(&balance, i, &wg)

		go func(b *int, n int, wg *sync.WaitGroup) {
			m.Lock()
			defer m.Unlock()
			*b -= n
			wg.Done()
		}(&balance, i, &wg)
	}

	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
