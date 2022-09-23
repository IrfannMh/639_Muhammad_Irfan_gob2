package main

import (
	"fmt"
	"strings"
)

func main() {
	//array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var names = [3]string{"irfan", "fahmi", "arya"}
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", names)
	fmt.Println("\n____________________")

	//modify array
	names[0] = "ucok"
	fmt.Printf("%#v\n", names)
	fmt.Println("\n____________________")

	//loopin array element
	//cara pertama
	for i, v := range names {
		fmt.Printf("Index : %d, Value : %s", i, v)
	}

	fmt.Printf(strings.Repeat("#", 25))
	//cara kedua
	for i := 0; i < len(names); i++ {
		fmt.Printf("index : %d, value: %s\n", i, names[i])
	}

	fmt.Printf(strings.Repeat("#", 25))

	//multidimensi array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		fmt.Printf("\n%d", arr)
		for _, value := range arr {
			fmt.Printf("\n%d", value)
		}
		fmt.Println()
	}

	var fruits1 = []string{"app", "bnn", "mng", "durian", "nanas"}
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v\n", fruits1)
}
