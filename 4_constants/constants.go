package main

import "fmt"

const age = 300

// name := "hello" // no
var name string = "golang"

func main() {
	// const name string = "golang"
	// name = "javascript" // not allowed

	// const age = 30

	// fmt.Println(age)

	const (
		port = 5000 
		host = "localhost"
	)

	// port = 5500 // error // the value assigned only when initialized

	fmt.Println(port, host)

}