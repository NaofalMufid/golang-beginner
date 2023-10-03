package main

import "fmt"

func sum(x []int, c chan int) {
	sum := 0
	for _, y := range x {
		sum += y
	}
	c <- sum
}

func main() {
	s := []int{1, 4, 8, -5, 9, 3}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
