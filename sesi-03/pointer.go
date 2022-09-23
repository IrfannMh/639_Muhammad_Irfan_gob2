package main

import (
	"fmt"
)

// func main(){
// 	var firstNumber int = 4
// 	var secondNumber *int

// 	fmt.Println("FirstNumber (value) :", firstNumber)
// 	fmt.Println("FirstNumber (memori address) :", &firstNumber)

// 	fmt.Println("SecondNumber (value) :", secondNumber)
// 	fmt.Println("SecondNumber (memori address) :", &secondNumber)
// }

// func main(){
// 	var firstPerson string = "John"
// 	var secondPerson *string = &firstPerson

// 	fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memori address) :", &firstPerson)
// 	fmt.Println("secondPerson (value) :", secondPerson)
// 	fmt.Println("secondPerson (memori address) :", &secondPerson)

// 	fmt.Println("\n", strings.Repeat("#", 30), "\n")

// 	*secondPerson = "Doe"

// 	fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memori address) :", &firstPerson)
// 	fmt.Println("secondPerson (value) :", secondPerson)
// 	fmt.Println("secondPerson (memori address) :", &secondPerson)
// }

// pointer as parameter
func main() {
	var a int = 10

	fmt.Println("Before: ", a)

	changeValue(&a)

	fmt.Println("After: ", a)
}

func changeValue(number *int) {
	*number = 20
}
