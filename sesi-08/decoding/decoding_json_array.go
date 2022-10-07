package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json"age"`
}

func main() {
	var jsonString = `[
		{
			"full_name":"Airel Jordan",
			"email":"airell@gmail.com",
			"age":23
		},
		{
			"full_name":"Ananda Jordan",
			"email":"ananda@gmail.com",
			"age":25
		},
		{
			"full_name":"Anhar Jordan",
			"email":"anhar@gmail.com",
			"age":23
		}
	]`
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

}
