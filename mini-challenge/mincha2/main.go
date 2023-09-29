package main

import "fmt"

func main() {
	merk := "toyota supra"
	// string per karakter
	for _, v := range merk {
		fmt.Printf("%c\n", v)
	}

	// hitung kemunculan jumlah perkarakter
	characterCount := make(map[rune]int)
	for _, v := range merk {
		characterCount[v]++
	}

	// tampilkan jumlah perkarakter
	for char, count := range characterCount {
		fmt.Printf("%c:%d ", char, count)
	}
	fmt.Println()
}
