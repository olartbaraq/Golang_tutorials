package main

import (
	"fmt"
)

//var x = 8

func main() {
	names := [5]string{"Ahmad", "Mubaraq", "Is'mail", "Yaqub", "Ibrahim"}
	friends := [2]string{"Ahmad", "Mubaraq"}

	_, _ = names, friends

	//Label

	// outer:
	// 	for index, name := range names {
	// 		for _, friend := range friends {
	// 			if name == friend {
	// 				result := []string{friend}
	// 				count := len(result)
	// 				fmt.Printf("we found %d friend %s at index %d \n", count, friend, index)
	// 				break outer
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("Ounter label done, waiting for instructions...")

	// Goto

	//switch cases

	language := "golang"

	switch language {
	case "python":
		fmt.Println("You are a python developer")
	case "golang", "Go":
		fmt.Printf("You are a golang developer")
	default:
		fmt.Println("Ogbeni, pick a language")
	}
}
