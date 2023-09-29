package model

import "strconv"

type User struct {
	ID     int
	Name   string
	Alamat string
	Posisi string
	Motto  string
}

func FindUser(users []User, input string) (User, bool) {
	// convert argumen input menjadi integer
	id, err := strconv.Atoi(input)
	// jika konversi berhasil
	if err == nil {
		// cari by id
		for _, user := range users {
			if user.ID == id {
				return user, true
			}
		}
	}

	// cari by nama
	for _, user := range users {
		if user.Name == input {
			return user, true
		}
	}
	// jika tidak  ditemukan balikkan user kosong dan kondisi false
	return User{}, false
}
