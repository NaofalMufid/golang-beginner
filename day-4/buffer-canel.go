package main

import "fmt"

func main() {
	c := make(chan bool, 5)
	c <- true

	fmt.Println("di tam pilkan")

	fmt.Println("-----Buffer channel ------")
	data := make(chan string, 2)
	data <- "Hey"
	data <- "You"

	firstPrint, secondPrint := <-data, <-data
	fmt.Println(firstPrint, secondPrint)
}
