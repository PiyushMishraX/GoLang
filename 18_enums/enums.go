package main

import "fmt"

// enumrated type

// custom type go
// type myType string


// we create num using custom type
// type OrderStatus int 

// const (
// 	Received OrderStatus = iota // iota is preeclared identifier for untyped integer
// 	Confirmed // 1 // auto increment // we can use string too
// 	Prepared // 2
// 	Delivered // 3 
// )

type OrderStatus string
const (
	Received OrderStatus = "received"
	Confirmed            = "confirmed"
	Prepared             = "prepared"
	Delivered  			 = "delivered"
)


// func changeOrderStatus(status string) {
// 	fmt.Println("Changing order status to", status)
// }

func changeOrderStatus1(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
//  changeOrderStatus("received")
//  changeOrderStatus("confirmed")
 // here might have typo , and there we have to use those for multiple time
 // we have to write everywhere for normal name change too
 // order status type can be vast but limited 
 // go do not have inbuilt enum , we implement using const and type

 // integer
//  changeOrderStatus1(Received)
//  changeOrderStatus1(Confirmed)
//  changeOrderStatus1(Prepared)

// string
 changeOrderStatus1(Received)
 changeOrderStatus1(Confirmed)
 changeOrderStatus1(Prepared)
 
 

}