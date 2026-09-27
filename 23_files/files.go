package main

import (
	// "fmt"
	// "fmt"
	"bufio"
	"fmt"
	// "io"
	"os"
	// "strings"
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


	

	//  READ FILE

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

	// data, err := os.ReadFile("example.txt") 
	// if err != nil { // can abstract this in a function
	// 	panic(err)
	// }

	// fmt.Println(string(data))
	// // we should not use this fn all the time
	// // the ReadFile function loads the file at once in memory the file size when small will be good reading through this mehtod 
	// // but bigger files such as video(GB's of data) should bot load like this , it isn't viable resouces might be less than this
	// // we can use streaming like node js in go too // will see in writing files time 





	// READ FOLDERS 

	// // dir, err := os.Open(".") // . current folder
	// dir, err := os.Open("../") // root 
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// // fileInfo, err := dir.ReadDir(1) // read info of DIr and returns slice / list of info 
	// // (1) // 1 value readed , 2 , 3 have to inputed

	// // fileInfo, err := dir.ReadDir(2) 
	// // fileInfo, err := dir.ReadDir(3) // still 2
	// fileInfo, err := dir.ReadDir(-1) // all files 

	// for _, fi := range fileInfo {
	// 	// fmt.Println(fi.Name())
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }


	// // CREATE A FILE
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go ")
	// f.WriteString("very nice language") // append mode 
	// // everytime the above code execute the file is recreated

	// // replace content

	// content, err := os.ReadFile("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }


	// newContent := strings.ReplaceAll(string(content), "very", "IT IS CHANGED")

	// os.WriteFile("example2.txt", []byte(newContent), 0644)



	// method 2 - to enter data
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// // file is nothing but byte data 

	// bytes := [] byte("golang HELLO") // slice of byte
	// f.Write(bytes)



	// reading from a file and adding in another file , 
	// in straming fashion ( not loadng all data in memory)
	// READ AND WRITE to another file ( streaming fashion)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example3.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	// streaming fashion through buf io mehtod
	
	reader := bufio.NewReader(sourceFile) //data from reader // newreader returns a reader whose buffer deafult size is 4096 byte
	writer := bufio.NewWriter(destFile) // writer


	for { // read byte by byte and write byte by byte in infi loop till EOF

		b, err := reader.ReadByte()
		if err != nil {
			//  checking extra err , end of file error

			if err.Error() != "EOF" { // returns error in string format 
				panic(err)  // if not EOF panic else break from infi loop 
			}

			// panic(err) // we don't just panic all the time , just writing for the time

			break // break when EOF error occurs so the infi loop ends
		}

		e := writer.WriteByte(b) // pass byte "b" // the err have to used like x, err to not throw error
		if e != nil {
			panic(e)
		}
		
	}

	// flush write for any remainign data at last
	writer.Flush()

	fmt.Println("Written to new file is successful")
	// the above is straming fashion file copying
	// but if we only need to copy we can use copy function , from source to destination file

	// _, err = io.Copy(destFile, sourceFile)
	// if err != nil {
	// 	panic(err)
	// }
	// destFile.Sync()









}