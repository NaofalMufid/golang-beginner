package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("goroutine pertama")
	// go func() {
	// 	fmt.Println("goroutine lainnya")
	// }()

	// fmt.Println("goroutine kedua")
	// go func() {
	// 	fmt.Println("goroutine dan lainnya")
	// }()
	// time.Sleep(time.Second)

	fmt.Println("goroutine ketiga")

	go firstPrint()
	go secondPrint()

	time.Sleep(time.Millisecond * 5)
}

func firstPrint() {
	fmt.Println("goroutine lain lain")
}

func secondPrint() {
	fmt.Println("goroutine lain pula")
}
