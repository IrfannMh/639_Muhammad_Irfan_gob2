package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	name     string
	age      int
	division string
}

type Employee1 struct {
	division string
	person   Person
}

func main() {
	// Giving value to struct
	var employee Employee

	employee.name = "Airell"

	employee.age = 23

	employee.division = "Curriculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	// Initializing struct
	var employee1 = Employee{}
	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curriculum Developer"

	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	// Pointer to a struct
	var employee3 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee3 name:", employee3.name)

	fmt.Println(strings.Repeat("#", 20))

	employee3.name = "Ananda"

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee3 name:", employee3.name)

	// Embedded struct
	var employee4 = Employee1{}

	employee4.person.name = "Airell"
	employee4.person.age = 23
	employee4.division = "Curriculum Developer"

	fmt.Printf("%+v\n", employee4)

	// Anonymous struct
	// Anonymous struct tanpa pengisian field
	var employee5 = struct {
		person   Person
		division string
	}{}

	employee5.person = Person{name: "Airell", age: 23}
	employee5.division = "Curriculum developer"

	// Anonymous struct dengan pengisian field
	var employee6 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Airell", age: 23},
		division: "Curriculum developer",
	}

	fmt.Printf("Employee5: %+v\n", employee5)
	fmt.Printf("Employee6: %+v\n", employee6)

	// Slice of struct
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	// Slice of anonymous struct
	var people2 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range people2 {
		fmt.Printf("%+v\n", v)
	}
}
