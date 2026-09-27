package main

import (
	// "fmt"
	"fmt"
	"os"
)

// working with files
// read , writ etc

func main () {
	// Info Get

	// // os.Open("example.txt") // name / path
	// f, err := os.Open("example.txt")  // rturn file object pointer and error // in go err handdling is very imp 
	// // go we return error not throw it

	// // er2 := 0 
	// // er2 := 0
	// // fmt.Println(er2) 
	// // can redeclare a variable named err in Go, but only within 
	// // the same block using multi-variable short variable declarations (:=).

	
	// if err != nil {
	// 	// we can
	// 	// log the error
	// 	// panic the err // panic(err)

	// 	panic(err) // see in adv error handling
	// }

	// fileInfo, err := f.Stat() // file info, err 
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name()) // we used relative path from the folder so the running hav eto be from there


	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size(bytes):", fileInfo.Size()) 
	// fmt.Println("file permission:", fileInfo.Mode()) 
	// fmt.Println("file modified at:", fileInfo.ModTime()) 


	//  read file

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close() // runs after main end // file close

	// // data read stored in buffer
	// // buf := make([]byte, 10) // bytes ka array // size can be size of file // have to create stack
	// buf := make([]byte, 12) 

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// // fmt.Println is the standard, production-safe function for formatted output, while println is a built-in debugging tool with specific limitations. 

	// // println("data", d, buf) // d is no of bytes

	// //  to read convert to string

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i])) // last two char not seen when buf = 10
	// }


	// read file easier method

	data, err := os.ReadFile("example.txt") 
	if err != nil { // can abstract this in a function
		panic(err)
	}

	fmt.Println(string(data))
	// we should not use this fn all the time
	// the ReadFile function loads the file at once in memory the file size when small will be good reading through this mehtod 
	// but bigger files such as video(GB's of data) should bot load like this , it isn't viable resouces might be less than this
	// we can use streaming like node js in go too // will see in writing files time 




}