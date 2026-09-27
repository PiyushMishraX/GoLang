package main

import (
	"fmt"
	"time"
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
// goroutine synchronizer
// func task(done chan bool) {

// 	defer func () { done <- true} ()  // runs after fn running ends
// 	fmt.Println("Processing...")
// 	// done<-true // can't reach when errors in b/w
// }



// func emailSender( emailChan chan string, done chan bool) {

// 	defer func() { done <- true}()

// 	for email := range emailChan{
// 		fmt.Println("Sending email to", email)

// 		time.Sleep(time.Second) // simulate low time processing of emails
// 	}
// }
// background workers etc are created through channels // above is a worker


//  type safety increament 
// func emailSender( emailChan chan string, done chan bool) {
func emailSender( emailChan <-chan string, done chan<- bool) {
	/// recieve only <- , <- send only 

	defer func() { done <- true}()

	// <-done
	// emailChan <- "email"
	// the abvoe tasks can be done but should not be  // so creating receive only and send only

	for email := range emailChan{
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second) 
	}
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

	// done := make(chan  bool)
	// go task(done)

	// <- done // block  //till someone sends data
	// single goroutine --> use channel
	// multiple --> use wg 



	//  problem channel 
	// send and receive is blocking 
	// but in such as queue system they are very slow
	// unbuffered blocking an process
	// soln -> buffer channel --> can send limited ammount of data without blocking 

	// ex --> email queue system
	// emailChan := make(chan string, 100) // struct
	// // 100 --> buffer size
	// done := make(chan bool)

	// go emailSender(emailChan, done) // trigger fn // done synchronization the two go routines 

	// // for i := 0; i < 100; i++ {
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// } // non blocking // deadlock at the end // we have to close fn
	

	// fmt.Println("Done sending")

	// close(emailChan) // the range of worker is set to 10  // this is important

	// // emailChan <- "1@example.com"
	// // emailChan <- "2@example.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan) // no deadlock

	// <-done // go routine block // range have infinite lock in the worker have so the blocking is for always that why deadlock




	// Multiple channel data sending back in channel
	chan1 := make(chan int)
	chan2 := make(chan string)

	// data send
	go func() {
		chan1 <- 10
	} () // channel not passed but closure used to run // value from parent used and saved

	go func() {
		chan2 <- "pong"
	} ()

	// data recive // for loop + select case
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}



}