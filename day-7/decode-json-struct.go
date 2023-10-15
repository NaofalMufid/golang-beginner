package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `{
		"full_name" : "Black Panda",
		"email" : "pandablack@mail.com",
		"age" : 21
	}`

	var result Student

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Fullname :", result.FullName)
	fmt.Println("Email :", result.Email)
	fmt.Println("Age :", result.Age)
}
