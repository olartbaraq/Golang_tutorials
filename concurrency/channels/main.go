package main

import "fmt"

func main() {
	//**CHANNELS**//

	// Channels are used to communicate in between goroutines. The data being sent or received
	// from a channel must be of the same type with the data specified when declaring the channel

	//**DECLARING A CHANNEL**//
	var ch chan int // a variable ch of type chan int.
	// we wan to communiate with a channel of type int

	// Value of an uninitialize channel is <nil>
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	// USing the short declaration

	newCh := make(chan int)

	// There are two main operations to perform on a channel.
	// SEND AND RECEIVE

	//SEND
	newCh <- 10
	fmt.Println(newCh)

	// RECEIVE
	num := <-newCh
	_ = num

	//OTHER OPERATIONS ARE
	// CLOSE
	close(newCh)

	// MAKING A UNIDRIECTIONAL CHANNEL

	c1 := make(<-chan int) // only for receiving
	_ = c1

	c2 := make(chan<- string) // only for sending
	_ = c2
}
