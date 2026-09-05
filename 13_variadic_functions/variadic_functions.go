package main

import "fmt"

// a variadic functions 
// mutiple paramters
func sum(nums ...int) int { // ...<- in the left of type
	total := 0

	for num := range nums {
		total = total + num

	}

	return total
}

// func sum(nums ...interface{}) int { // any type
// // func sum(nums ...any) int {
// 	total := 0

// 	for num := range nums {
// 		total = total + num
// 	}
// 	return total
// }

func main() {
	// variadic functions

	// fmt.Println(1, 2, 3 , 4, 5, 86, "hello") // can have n number of parameterss just like fmt.Println

	// result := sum(2, 4, 6,3, 63, 5)// integers only
	// fmt.Println(result)

	nums := []int{3,4,5,6}
	result := sum(nums...) // just like spread operator // in the right of element
	fmt.Println(result)




}