package main

import "fmt"

// range --> iterating over data structures
func main(){
	// nums := []int {6, 7, 8}

	// with for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// with range
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }


	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	// range in map
	// m := map[string]string{"fname": "marine", "lname": "diver"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// m := map[string]string{"fname": "marine", "lname": "diver"}
	// for k := range m { // only key
	// 	fmt.Println(k )
	// }


	// range in string
	// ascii --> only till 255 ( 1 byte ), unicode is beyond it ( 1 byte , 2 byte)
	for i, c := range "golang" { // i is the starting byte not index so for 2 byte ones the indexes would be 0 2 4 6 etc
		// fmt.Println(i, c) // c --> unicode code 
		fmt.Println(i, c , string(c))
	}



}