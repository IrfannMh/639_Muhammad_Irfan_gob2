package main

import "fmt"

func main() {
	//array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"irfan", "fahmi", "arya"}
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
	fmt.Println("\n____________________")

	//modify array
	strings[0] = "ucok"
	fmt.Printf("%#v\n", strings)
	fmt.Println("\n____________________")

}
