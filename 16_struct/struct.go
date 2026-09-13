package main

import (
	"fmt"
	"time"
)

// "fmt"
// "time"

// go do not have classes so we use structure

// order struct
// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createAt time.Time // nanosecond precision
// } // used via instance creation

// func main(){
// 	// var order order =
//
// 	myOrder := order{
// 		id: "11",
// 		amount: 500.00,
// 		status: "received",
// 	}
//
// 	// fmt.Println("Order struct", myOrder)
//
// 	myOrder.createAt = time.Now()
// 	// fmt.Println("Order struct", myOrder)
//
// 	// fmt.Println(myOrder.status)
//
// 	myOrder2 := order{
// 		id: "12",
// 		status: "delivered",
// 		amount: 100,
// 		createAt: time.Now(),
// 	}
//
// 	myOrder.status = "paid"
//
// 	fmt.Println("Order struct", myOrder)
// 	fmt.Println("Order struct", myOrder2)
//
// }

/* Adding methods in struct --> so they work as object */

type order struct {
	id string
	amount float32
	status string
	createAt time.Time 
}

// receiver type
// func (o order) statusChange(status string) {
func (o *order) statusChange(status string) {
	// function attached to struct

	o.status = status // struct derefrence automatically
}

// for modification using pointer is required but getting will work fine for both

func (o order) getAmount() float32 {
	return o.amount
}


func main() {

	// myOrder := order{
	// 	id: "12",
	// 	status: "delivered",
	// 	amount: 100,
	// 	createAt: time.Now(),
	// }

	// // myOrder.statusChange("confirmed")
	// // fmt.Println(myOrder)

	// fmt.Println(myOrder.getAmount())

	// if we don't set any field , 
	// default will be zero value
	// int -> 0 , float -> 0 , string -> "", bool-> false
	myOrder := order{
		// id: "12",
		// status: "delivered",
		// amount: 100,
		// createAt: time.Now(),
	}

	fmt.Println(myOrder)

}