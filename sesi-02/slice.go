package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[3] = "mango"

	//append function
	fruits = append(fruits, "apple", "banana", "mango")

	fmt.Printf("%#v", fruits)

	//append dengan ellipsis (...)
	// var fruits1 = []string{"apple","banana","mango"}
	// var fruits2 = []string{"durian","pineapple","starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	//copy elements
	// var fruits1 = []string{"apple","banana","mango"}
	// var fruits2 = []string{"durian","pineapple","starfruit"}
	// nn := copy(fruits1,fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copy elements =>", nn)

	//slicing
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits1[:]
	fmt.Printf("%#v\n", fruits5)

	//slicing & append
	var fruits6 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v", fruits6)

	//backing array
	var fruits7 = fruits1[2:4]
	fruits7[0] = "rambutan"

	fmt.Println("fruits backing array =>", fruits7)

	//cap array
	fmt.Println("Fruits cap:", cap(fruits1))
	fmt.Println("Fruits len:", len(fruits1))

	fmt.Println(strings.Repeat("#", 30))
	var fruits8 = fruits1[0:3]

	fmt.Println("Fruits cap:", cap(fruits8))
	fmt.Println("Fruits len:", len(fruits8))

	fmt.Println(strings.Repeat("#", 30))
	var fruits9 = fruits1[1:]

	fmt.Println("Fruits cap:", cap(fruits9))
	fmt.Println("Fruits len:", len(fruits9))

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("new cars:", newCars)

}
