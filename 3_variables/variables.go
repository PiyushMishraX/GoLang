package main

import "fmt"

func main() {
	var name string = "golang" // in golang any created variable must be used or deleted

	fmt.Println(name)
	fmt.Println((name))

	var name2 = "golang" // golang autmatically infer type
	fmt.Println(name2)

	var isAdult = true // bool
	fmt.Println(isAdult)

	var age int = 30// go internally optimize size 
	// var age int // auto initialize with 0
	fmt.Println(age)


	// shorthand 
	name3 := "golang"
	fmt.Println(name3)

	// situation where writing type is must
	// can not writ --> // var name4 // if no initialization
	var name4 string
	fmt.Println(name4)

	// var price float32 = 50.5
	// var price = 50.5
	price := 50.5
	price = 2
	fmt.Println(price)


	




}