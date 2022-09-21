package main

import "fmt"

func main() {
	//looping
	//cara pertama
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("\n____________________")

	//cara kedua

	var i = 0
	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}
	fmt.Println("\n____________________")

	//cara ketiga
	i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 3 {
			break
		}
	}
	fmt.Println("\n____________________")

	//break & continue
	for y := 1; y <= 10; y++ {
		if y%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", y)
	}
	fmt.Println("\n____________________")

	//nested & label
outerLoop:
	for i = 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, "")
		}
		fmt.Print("\n")
	}
	fmt.Println("\n____________________")
}
