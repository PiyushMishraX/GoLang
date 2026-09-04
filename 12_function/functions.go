package main


// go can return multiple return values
// the errors can be returned too

// argument and return value
// func add(a int, b int) int { // return value
// func add(a, b int) int {
// 	return a + b
// }

// 
// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true // multiple return values with the type written above in ( )
// }

func processIt(fn func(a int) int) {
	fn(1)
}

// func processIt() func(a int) int {
// 	return func(a int) int {
// 		return 4
// 	}
// }

// main funtions --> main function is entry point of code 
func main() {
	// result := add(3, 5)
	// fmt.Println(result)

	// lang1, lang2, lang3-k := getLanguages()
	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	// fn := processIt()
	// fn(6)

}