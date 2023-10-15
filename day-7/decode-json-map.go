package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{
		"full_name" : "Black Panda",
		"email" : "pandablack@mail.com",
		"age" : 21
	}`

	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Fullname :", result["full_name"])
	fmt.Println("Email :", result["email"])
	fmt.Println("Age :", result["age"])
}
