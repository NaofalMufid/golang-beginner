package main

import (
	"fmt"
	"sync"
)

var (
	interface1 interface{}
	interface2 interface{}
	wg         sync.WaitGroup
)

func main() {
	interface1 = []string{"test1", "test2", "test3"}
	interface2 = []string{"jalan1", "jalan2", "jalan3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printProcess(i, &wg, interface1)
		go printProcess(i, &wg, interface2)
	}

	wg.Wait()
}

func printProcess(i int, wg *sync.WaitGroup, interfaceData interface{}) {
	fmt.Println(interfaceData, i)
	wg.Done()
}
