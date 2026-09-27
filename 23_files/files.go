package main

import (
	"fmt"
	"os"
)

// working with files
// read , writ etc

func main () {
	// os.Open("example.txt") // name / path
	f, err := os.Open("example.txt")  // rturn file object pointer and error // in go err handdling is very imp 
	// go we return error not throw it

	// er2 := 0 
	// er2 := 0
	// fmt.Println(er2) 
	// can redeclare a variable named err in Go, but only within 
	// the same block using multi-variable short variable declarations (:=).

	
	if err != nil {
		// we can
		// log the error
		// panic the err // panic(err)

		panic(err) // see in adv error handling
	}

	fileInfo, err := f.Stat() // file info, err 
	if err != nil {
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name()) // we used relative path from the folder so the running hav eto be from there


	fmt.Println("file or folder:", fileInfo.IsDir())
	fmt.Println("file size(bytes):", fileInfo.Size()) 
	fmt.Println("file permission:", fileInfo.Mode()) 
	fmt.Println("file modified at:", fileInfo.ModTime()) 


}