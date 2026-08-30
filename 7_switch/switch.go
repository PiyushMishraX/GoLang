package main

import "fmt"

func main() {
	// simple switch -- in go it is not required to
	// use break in switch explicitly
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// both abive and below are same

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	//      break
	// case 2:
	// 	fmt.Println("two")
	//      break
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// type switch -- > to tell type of variable
	whoAmI := func(i interface{}) {

		// switch t = i.(type) {// cannot use like this
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			// fmt.Println("other",t ) // cannot ise like this
			fmt.Println("other")
		}
	}

	whoAmI(55)
	whoAmI("go")
	whoAmI(true)
	whoAmI(whoAmI)

}
