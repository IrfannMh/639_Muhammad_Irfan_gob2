package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://developer.com:80/hello?name=airell&age=23#hello"
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url : %s\n", urlString)  //get urlString
	fmt.Printf("url : %s\n", u.String()) //get urlParse

	fmt.Printf("protocol : %s\n", u.Scheme)
	fmt.Printf("host : %s\n", u.Host)
	fmt.Printf("path : %s\n", u.Path)

	fmt.Printf("port : %s\n", u.Port())
	// fmt.Printf("fragment : %s\n", u.Opaque)
	fmt.Printf("fragment : %s\n", u.Fragment)
	fmt.Printf("raw Query (isi parameter) : %s\n", u.RawQuery)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name : %s, age : %s\n", name, age)

}
