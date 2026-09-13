package main

import "fmt"

func counter() func() int { // receives nothig returns a function returning an interger
	var count int = 0

	return func() int { // a closure
		count += 1
		return count
	}

	// whenever a function executes , it goes in cal stack 
	// and when the excution ends it gets removed from the call stack
	// thus variables are removed

	// but in a closure outer scope variable is used then the 
	// vareiable is always present inside that closure

}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}