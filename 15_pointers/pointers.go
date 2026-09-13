package main

import "fmt"

// by value --> distict copy passed 
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

// by reference
func changeNum2(num *int) {
	// num = 5 // pointer deference than change, pointer ke andar ki value change
	*num = 5 // num is pointer 
	fmt.Println("In changeNum", *num)
}

// pointers // memory location of a variable
func main() {
	num := 1

	// changeNum(num) // here only the variables copy goes not the real num
	// fmt.Println("After changeNum in main", num)

	// changing the main variables --> passing the reference of variable
	// fmt.Println("Memory address", &num)

	changeNum2(&num) // location passed so source variable changes
	fmt.Println("After changeNum in main", num) 


}