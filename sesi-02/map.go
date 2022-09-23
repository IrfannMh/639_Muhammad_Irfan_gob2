package main

import "fmt"

func main(){
	var person map[string]string //deklarasi

	person = map[string]string{} //inisialisasi

	person["name"] = "Airell"

	person["age"] = "23"

	person["address"] = "Jalan Sudirman"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])
	

	//cara ke2
	var person1 = map[string]string{
		"name":	"Airell",
		"age":	"23",
		"address":	"Jl Mangga"
	}

	fmt.Println("name:", person1["name"])
	fmt.Println("age:", person1["age"])
	fmt.Println("address:", person1["address"])


	//looping with map
	for key,value := range person1{
		fmt.Println(key, ":", value)
	}

	//deleting value
	fmt.Println("Before deleting:", person1)

	delete(person1, "age")

	fmt.Println("After deleting:", person1)

	//detecting value
	value, exist := person["age"]

	if exist{
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person, "age")
	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}


	//combining slice and map

	var people = []map[string]string{
		{"name":"Airell", "age":"23"},
		{"name":"Nanda", "age":"23"},
		{"name":"Mailo", "age":"15"},
	}

	for i, person := range people{
		fmt.Printf("Index: %d, name: %s, age: %s", i, person["name"],person["age"])
	}
}