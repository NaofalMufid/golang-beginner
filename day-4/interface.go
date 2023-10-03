package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter, jari_jari float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l lingkaran) luas() float64 {
	return math.Phi * math.Pow(l.jari_jari, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func (l lingkaran) keliling() float64 {
	return math.Phi * l.diameter
}

func main() {
	var lingkaran1 hitung = lingkaran{14, 7}
	var persegi1 hitung = persegi{4}

	fmt.Printf("Tipe data lingkaran %T\n", lingkaran1)
	fmt.Printf("Tipe data persegi %T\n", persegi1)
}
