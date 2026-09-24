package main

import (
	"fmt"
	// "math/rand"
	// "time"
	// "time"
)

// communication between goroutines happens through channels


// sending
// func processNum(numChan chan int) {
// 	// fmt.Println("Processing number", <-numChan)

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		// fmt.Println("Processing number", <-numChan)
// 		time.Sleep(time.Second * 1)

// 		// queue system can be implemented through this
// 	}
// }


// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }


// wg work ( holding main function ) using channels
func task(done chan bool) {

	defer func () { done <- true} ()  // runs after fn running ends

	fmt.Println("Processing...")

	// done<-true // can't reach when errors in b/w
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




	// numChan := make(chan int)

	// main is seperate go routine and precessNum is other one

	// go processNum(numChan) 
	

	// time.Sleep(time.Second * 2)
	// fmt.Println("1")

	// numChan <- 5 // number send to 1 goroutine to other through channel
	
	// fmt.Println("2")
	// time.Sleep(time.Second * 2)



	// we use channels likek queue in out channel
	// numChan := make(chan int)
	// go processNum(numChan)
	// // nunCHan <-5
	// for { // infinite loop no need for sleep
	// 	numChan <- rand.Intn(100) // rand num bw 0 and 100
	// }

	// // time.Sleep(time.Second * 2)




	// receiving
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result // data receive // no nned for time sleep because it is blocking ( means this holds the function till the running completes)

	// fmt.Println(res)


	// wg alternative

	done := make(chan  bool)
	go task(done)

	<- done // block  //till someone sends data
	



}