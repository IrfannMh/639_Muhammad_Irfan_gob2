package main

import "fmt"

//buat login
//bandingkan dengan username & pass
//username -> string
//pass -> ****
// if username && password  == true
// username dan password benar
// else
// username/password salah

func main() {
	
	username1 := "irfan"
	password1 := "golang"
	
	_, _ = username1, password1
	
	
	checkLogin(username1, password1)
	
}
func asteriskWord(w) {
	return '*'.repeat(w.length)
}
func checkLogin(username, password string) {
	var user, pass string
	fmt.Scanln(&username)
	fmt.Scanln(&pass)
	if user == username && password == pass {
		fmt.Println("Username dan password benar")
	} else {
		fmt.Println("username / password salah")
	}
}
