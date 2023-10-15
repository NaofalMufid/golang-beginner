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
	var jsonString = `[
		{
			"full_name" : "Black Panda",
			"email" : "pandablack@mail.com",
			"age" : 21
		},
		{
			"full_name" : "Black Rabbit",
			"email" : "rabibtblack@mail.com",
			"age" : 23
		}
	]`

	var result []Student

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("index %d : %+v\n", i+1, v)
	}
}
