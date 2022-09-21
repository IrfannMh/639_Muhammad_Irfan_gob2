package main

import "fmt"

func main() {
	//komentar 1 bari
	fmt.Println("Hello")

	fmt.Println("Scalable", "web", "Service", "Golang")
	/*
		ini komentar
		lebih dari 1 baris
	*/
	fmt.Println("Muhammad", "Irfan\n")
	fmt.Println("Muhammad", " Irfan\n")
	fmt.Println("Muhammad", " ", "Irfan\n")
	fmt.Println("___________________")

	//Belajar variabel golang
	//variabel dengan data type

	var namaPanjang string = "Irfan"
	var umur int = 22

	namaPanjang = "Muhammad Irfan"

	fmt.Println("Nama Saya ", namaPanjang)
	fmt.Println("Umur Saya ", umur)
	fmt.Println("___________________")

	//tanpa data type
	var nama = "irfan"
	var age = 23
	fmt.Println("Nama Saya ", nama)
	fmt.Println("Umur Saya ", age)
	fmt.Println("___________________")

	//short declarations
	name := "M Irfan"
	age1 := 23
	fmt.Println("%T, %T", name, age1)
	fmt.Println("___________________")

	//multiple declaration
	var s1, s2, s3 string = "satu", "dua", "tiga"
	var i1, i2, i3 int
	i1, i2, i3 = 1, 2, 3

	fmt.Println(s1, s2, s3)
	fmt.Println(i1, i2, i3)
	fmt.Println("___________________")

	//underscore variabel
	var variabelSatu string
	var satu, dua, usia, alamat, oneName = "satu", "dua", 24, "Gg Mangga", "Irfan"
	_, _, _ = variabelSatu, satu, dua
	fmt.Printf("Variabel satu bertipe => %T", satu)
	fmt.Printf("\nHai nama saya %s, saya berusia %d, saya tinggal di %s", oneName, usia, alamat)
	fmt.Println("___________________")

	//tipe data
	var first uint8 = 80
	var second int8 = -1

	fmt.Printf("tipe data variabel first : %T\n", first)
	fmt.Printf("tipe data variabel second : %T\n", second)
	fmt.Println("___________________")

	//tipe data float
	var decimalNumber float32 = 3.12

	fmt.Printf("desimal number : %f\n", decimalNumber)
	fmt.Printf("desimal number 3 angka dibelakang koma : %.3f\n", decimalNumber)
	fmt.Println("\n___________________")

	//boolean
	var condition bool = true

	fmt.Printf("is permitted? %t\n", condition)
	fmt.Println("\n___________________")

	//string
	var message string = "Hallo"
	fmt.Print(message)
	fmt.Println("\n___________________")

	/*	Zero value pada masing masing tipe data
		Zero value dari string adalah "" (string kosong).
		Zero value dari bool adalah false.
		Zero value dari tipe numerik non-desimal adalah 0.
		Zero value dari tipe numerik desimal adalah 0.0.
	*/

	//constant
	const full_name string = "Muhammad Irfan"
	fmt.Printf("Hallo %s", full_name)
	fmt.Println("\n___________________")

	//operator
	var value = 2 + 2*3
	fmt.Println(value)

	//operator relational
	var first_cond bool = 2 > 3
	var sec_cond bool = "joey" == "Joey"
	var third_cond bool = 10 != 2.3
	var fourth_cond bool = 11 <= 11

	fmt.Println("first condition :", first_cond)
	fmt.Println("second condition :", sec_cond)
	fmt.Println("third condition :", third_cond)
	fmt.Println("fourth condition :", fourth_cond)
	fmt.Println("\n___________________")

	//logical operator
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%T) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%T) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t(%T) \n", wrongReverse)
	fmt.Println("\n___________________")

}
