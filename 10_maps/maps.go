package main

import "fmt"

// maps -> hash , object , dict
func main(){

	// creating map

	// m := make(map[string]string) // map[type_of_key]type_of_value

	// m["name"] = "golang"
	// m["area"] = "backend"

	// fmt.Println(m["name"], m["area"])
	// fmt.Println(m["phone"]) // if key value do not exist in the map then it returns zero value


	// m := make(map[string]int)
	// m["age"] = 30
	// // fmt.Println(m["phone"]) // 0
	// m["price"] = 50
	// fmt.Println(len(m))

	// delete(m, "price")
	// fmt.Println(m["price"])

	// clear(m)
	// fmt.Println(m)


	// other format to create
	// m := map[string]int{"price": 40, "phone": 3}
	// fmt.Println(m)



	m := map[string]int{"price": 40, "phone": 3}

	// _, ok := m["price"] // _ <-- do not have to be used

	// if ok { // element present
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	v, ok := m["price"] 

	fmt.Println(v) // value at that key
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}








}