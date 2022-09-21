package main

import "fmt"

func main() {

	//conditions
	var currentYear = 2022
	if age := currentYear - 1987; age > 17 {
		fmt.Printf("Kamu belom memiliki KTP ", age)
	} else {
		fmt.Println("Kamu sudah memiliki KTP")
	}
	fmt.Println("\n____________________")

	//switch
	var score = 8
	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}
	fmt.Println("\n____________________")

	//switch conditional & fallthrough
	var score1 = 6
	switch {
	case score1 == 8:
		fmt.Println("Perfect")
	case (score1 < 8) && (score1 > 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		fmt.Println("Study Harder")
		fmt.Println("You don't have a good score yet")
	}
	fmt.Println("\n____________________")

	//nested condition
	var score2 = 10
	if score2 > 7 {
		switch score2 {
		case 8:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score2 == 5 {
			fmt.Println("not bad")
		} else if score2 == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("you can do it")
			if score2 == 0 {
				fmt.Println("try harder")
			}
		}
	}

	fmt.Println("\n____________________")
}
