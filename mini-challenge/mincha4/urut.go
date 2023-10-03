package main

import (
	"fmt"
	"sync"
	"time"
)

type interface1 interface {
	Print()
}

type interface2 interface {
	Print()
}

type struct1 struct {
	name string
}

func (s struct1) Print() {
	fmt.Println(s.name)
}

type struct2 struct {
	name string
}

func (s struct2) Print() {
	fmt.Println(s.name)
}

func main() {
	// interface1
	s1 := struct1{
		name: "Pusing 1",
	}

	// interface2
	s2 := struct2{
		name: "Pusing 2",
	}

	// mutex
	mutex := &sync.Mutex{}

	// goroutine 1
	go func() {
		for i := 0; i < 4; i++ {
			// ambil kunci mutex
			mutex.Lock()

			// cetak data
			s1.Print()

			// lepas kunci mutex
			mutex.Unlock()
		}
	}()

	// goroutine 2
	go func() {
		for i := 0; i < 4; i++ {
			// ambil kunci mutex
			mutex.Lock()

			// cetak data
			s2.Print()

			// lepas kunci mutex
			mutex.Unlock()
		}
	}()

	// tunggu hingga goroutine selesai
	time.Sleep(time.Second * 1)
}
