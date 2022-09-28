package main

import "runtime/trace"

// func main() {
// 	var randomValues interface{}
// 	_ = randomValues

// 	randomValues = "Jalan Sudirman"

// 	randomValues = 20

// 	randomValues = true

// 	randomValues = []string{"Airell", "Nanda"}
// }

//type assertion
// func main(){
// 	var v interface{}
// 	v = 20
// 	// v = v*9 // error karena perkalian membutuhkan data yang konkrit atau asli
// 	if value, ok := v.(int); ok == true {
// 		v = value * 9
// 	}
// }

//map & slice empty interface
func main {
	rs := []interface{}{1,"Airell", true, 2, "Ananda", true}

	rm := map[string]interface{}{
		"Name": "Airell",
		"Status": true,
		"Age": 23,
	}

	_,_ = rs,rm
}