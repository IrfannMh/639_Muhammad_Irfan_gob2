package main

import (
	"fmt"
	"os"
)

// Defer#1
// func main() {
// 	defer fmt.Println("defer function starts to execute")
// 	fmt.Println("Hai everyone")
// 	fmt.Println("Welcome back to Go learning center")
// }

// Defer#2
// func main() {
// 	callDeferFunc()
// 	fmt.Println("Hai everyone!!")
// }

// func callDeferFunc() {
// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("Defer func starts to execute")
// }

// Exit
func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}
