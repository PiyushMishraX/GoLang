package main

import (
	"fmt"
	"time"
)

// communication between goroutines happens through channels



func processNum(numChan chan int) {
	fmt.Println("Processing number", <-numChan)

}


func main() {
	// messageChan := make(chan string) // string data to send

	// send data
	// messageChan <- "ping" // data send inside var
	// operation is blocking

	// recieve
	// <- messageChan
	// msg := <- messageChan 
	 
	// deadlock in above 

	// fmt.Println(msg)




	numChan := make(chan int)

	// main is seperate go routine and precessNum is other one

	go processNum(numChan) 

	numChan <- 5 // number send to 1 goroutine to other through channel

	time.Sleep(time.Second * 2)


}