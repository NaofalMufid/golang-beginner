package main

import "fmt"

func main() {
	// var cities = make([]string, 3)
	// cities[0] = "Surabaya"
	// cities[1] = "Solo"
	// cities[2] = "Samarinda"

	// cities = append(cities, "Palu", "Semarang")
	// cities2 := []string{"Denpasar", "Bandung"}

	// copyCities := copy(cities, cities2)
	// fmt.Println(cities)
	// fmt.Println(copyCities)

	// string per karakter
	merk := "toyota supra"
	for _, v := range merk {
		fmt.Printf("%c\n", v)
	}

	// hitung kemunculan jumlah perkarakter
	characterCount := make(map[rune]int)
	fmt.Println(characterCount)
	for _, v := range merk {
		characterCount[v]++
	}
	fmt.Println(characterCount)

	// tampilkan jumlah perkarakter
	for char, count := range characterCount {
		fmt.Printf("%c:%d ", char, count)
	}
	fmt.Println()
}
