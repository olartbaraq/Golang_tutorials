package main

import "fmt"

func findMissingNumber(n []int) int {

	nums := len(n) + 1
	mainResult := (nums * (nums - 1)) / 2
	sum := 0

	for _, v := range n {
		sum += v
		//fmt.Println(i, sum)
	}
	return mainResult - sum
}

func main() {

	c := make(chan int)

	go func(n []int, ch chan int) {

		nums := len(n) + 1
		mainResult := (nums * (nums - 1)) / 2
		sum := 0

		for _, v := range n {
			sum += v
			//fmt.Println(i, sum)
		}
		ch <- (mainResult - sum)
	}([]int{0, 1, 2, 3, 4, 5, 7, 8}, c)

	fmt.Println(findMissingNumber([]int{0, 1, 2, 4, 4}))
	fmt.Println(findMissingNumber([]int{3, 0, 1}))
	fmt.Println(findMissingNumber([]int{0, 1}))
	fmt.Println(findMissingNumber([]int{0, 1, 2, 3, 4, 6, 7}))

	fmt.Println(<-c)
	// const (
	// 	a = 6
	// 	b
	// 	c
	// )

	// fmt.Println(a, b, c)
}
