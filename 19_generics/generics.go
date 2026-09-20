package main

import "fmt"

// generics ( version >= 1.18)

// func printSlice( items []int) {
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printStringSlice( items []string) {
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }
// much duplication here , so using genrics to improce reusability 

// func printSlice[T interface{}]( items []T) { // convencion T
// func printSlice[T any]( items []T) { // convencion T
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }



// func printSlice[T int | string]( items []T) { // only acceptsnumber and string
// func printSlice[T int | string | bool]( items []T) { 
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }



//  LIFO
// type stack struct {
// 	elements []int// have to change here to use string etc
// }

// generic in struct
// type stack[T any] struct {
// 	elements []T
// }


// creating below more genric to hanle more type
// func printSlice[T comparable]( items []T) { // comparables types allowrd here 
func printSlice[T comparable, V string]( items []T, name V) { // comparables types allowrd here // mustiple generics passinf
	for _,item := range items {
		fmt.Println(item, name)
	}
}


func main() {
	// printSlice([]int{1, 2, 3})

	// nums := []int{1, 2, 3}
	// printSlice(nums)
	// names := []string{"golang", "typescript"}
	// printStringSlice(names)


	// nums := []int{1, 2, 3}
	// nums := []bool{true, false , true}
	// printSlice(nums)


	// myStack := stack {
	// 	elements: []int{1,2,3},
	// 	// elements: []string{"golang"},
	// }
	// fmt.Println(myStack)

	// myStack := stack[string] {
	// 	elements: []string{"golang"},
	// }
	// fmt.Println(myStack)

	vals := []bool{true, false , true}
	printSlice(vals, "hello")


}