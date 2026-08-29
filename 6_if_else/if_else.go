package main

import "fmt"

func main() {
	// age := 20

	// if age >= 21 {
	//	fmt.Println("Person is an Adult")
	// } else {
	//        fmt.Println("Person is not an Adult")
	// }

	// else if

	// age := 16
	// age = 10

	// if age >= 18 {
	// 	fmt.Println("Person is an Adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	//	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermission= true
	// var hasPermission = false

	// if role == "admin" || hasPermission {
	// if role == "admin" && hasPermission {
	//	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// variable creation with if block // we declare variable inside if construct
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is a teenager", age)
	} else {
		fmt.Println("Person is a kid", age)
	}
	// fmt.Println(age) // can not use outside if block
	
	// go does not have ternary operator
	// have to use normal if else 
	
	

}
