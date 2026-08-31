package main

import "fmt"

// slice --> dynamic array
// no need for length automatically increased if needed
// most used construct in go
// more useful methods
func main(){

	// declaration // create array and not mention size
	// var nums []int  // uninitialized slice is nil // nil values default

	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums)) // 0 for uninitialzed one


	// to not have nil in base
	// var nums = make([]int, 2) // initial size is 2
	// fmt.Println(nums)
	// fmt.Println(nums == nil)

	// var nums = make([]int, 2)
	// capacity --> maximum number of elements can fit // in slices it is initial allocated size
	// same as size for this(2) -> var nums = make([]int, 2)
	// fmt.Println(cap(nums))

	// var nums = make([]int, 2, 5)
	// fmt.Println(cap(nums)) // cap --> 5
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)
	// nums = append(nums, 11)
	// fmt.Println(nums)
	// fmt.Println(cap(nums)) // capacity doubles each time the limit is reached // it resizes

	// var nums = make([]int, 0, 5) // normal 0 so it would be empty slice
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	nums := []int{}
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))




}