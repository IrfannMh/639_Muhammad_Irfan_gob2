package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// city := "Jakarta"

	// for i:= 0; i<len(city); i++{
	// 	fmt.Printf("index: %d, byte: %d\n", i, city[i])
	// }

	//looping over string
	//byte by byte
	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println((string(city)))

	//rune by rune
	str1 := "makan"
	str2 := "manca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	for i, s := range str1 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}

}
