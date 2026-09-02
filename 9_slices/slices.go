package main

import (
	"fmt"
	// "slices"
)

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


	// nums := []int{}
	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 2 , 5)
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// // copy functions
	// var nums = make([]int, 0 , 5)
	// // var nums2 = make([]int, len(nums))
	// nums = append(nums, 2)
	
	// var nums2 = make([]int, len(nums))
	// //copy
	// fmt.Println(nums, nums2)
	
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)
	


	// slice operator
	// var nums = []int{1,2,3}

	// fmt.Println(nums[0:2]) // return elements from 0 till 2 elements are getten

	// fmt.Println(nums[:1]) // start tilll 1 element
	// fmt.Println(nums[1:])  // from this index to last


	// slice // package
	// var nums1 = []int{1,2,3}
	// var nums2 = []int{1,2,3,4}

	// fmt.Println(slices.Equal(nums1, nums2))


	
	// 2D SLICES

	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)



}