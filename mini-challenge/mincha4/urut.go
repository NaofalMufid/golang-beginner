package main

import (
	"fmt"
	"sync"
)

var (
	interface1 interface{}
	interface2 interface{}
	wg         sync.WaitGroup
	mt         sync.Mutex
)

func main() {
	interface1 = []string{"test1", "test2", "test3"}
	interface2 = []string{"jalan1", "jalan2", "jalan3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		mt.Lock()
		go printProcess(i, &wg, &mt, interface1)
		mt.Lock()
		go printProcess(i, &wg, &mt, interface2)
	}

	wg.Wait()
}

func printProcess(i int, wg *sync.WaitGroup, mt *sync.Mutex, interfaceData interface{}) {
	fmt.Println(interfaceData, i)
	mt.Unlock()
	wg.Done()
}
