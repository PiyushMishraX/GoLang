package main

import "fmt"

// in go only "for loop" is used /present for loop
func main() {

	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	//	println("1")
	// }

	//classic for loop
	// for i := 0; i < 3; i++ {
	// fmt.Println(i)
	// println(i) // println is for debugging not production
	// }

	// breaka nd continue
	// for i := 0; i <= 3; i++ {
	//	fmt.Println(i)
	//	break
	// }

	// for i := 0; i <= 3; i++ {
	//	if i == 2 {
	//		continue
	//	}
	//	fmt.Println(i)
	// }

	// range ( 1.22 )

	for i := range 11 { // range starting from 0 till 10 
		fmt.Println(i)
	}

}
