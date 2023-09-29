package main

import (
	"fmt"
	"mincha3/model"
	"os"
)

func main() {
	users := []model.User{
		{ID: 1, Name: "Yuno", Alamat: "Clover Kingdom", Posisi: "Lorem", Motto: "Ipsum"},
		{ID: 2, Name: "Mimosa", Alamat: "Clover Kingdom", Posisi: "Lorem", Motto: "Ipsum"},
		{ID: 3, Name: "Asta", Alamat: "Clover Kingdom", Posisi: "Lorem", Motto: "Ipsum"},
		{ID: 4, Name: "Yami", Alamat: "Clover Kingdom", Posisi: "Captain", Motto: "Ipsum"},
		{ID: 5, Name: "Charmy", Alamat: "Clover Kingdom", Posisi: "Lorem", Motto: "Ipsum"},
		{ID: 6, Name: "Mereleona", Alamat: "Clover Kingdom", Posisi: "Former Captain", Motto: "Ipsum"},
		{ID: 7, Name: "Charlotte", Alamat: "Clover Kingdom", Posisi: "Captain", Motto: "Ipsum"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Harap masukan Nomor Absen atau Nama Peserta: ")
		fmt.Println("Contoh 'go run main.go Sela' atau 'go run main.go 3' ")
		return
	}
	input := os.Args[1]

	user, found := model.FindUser(users, input)
	if found {
		fmt.Println("ID:", user.ID)
		fmt.Println("Nama:", user.Name)
		fmt.Println("Alamat:", user.Alamat)
		fmt.Println("Posisi:", user.Posisi)
		fmt.Println("Motto:", user.Motto)
	} else {
		fmt.Println("Data peserta tidak ditemukan")
	}
}
