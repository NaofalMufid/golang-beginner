package main

import (
	"fmt"
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
		name: "Mumet 1",
	}

	// interface2
	s2 := struct2{
		name: "Mumet 2",
	}

	// goroutine 1
	go func() {
		for i := 0; i < 4; i++ {
			s1.Print()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// goroutine 2
	go func() {
		for i := 0; i < 4; i++ {
			s2.Print()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// menunggu goroutine selesai
	time.Sleep(time.Second * 1)
}
