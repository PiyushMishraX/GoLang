package main

import (
	"fmt"
	"time"
	// "time"
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

// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createAt time.Time
// }

// receiver type
// func (o order) statusChange(status string) {
// func (o *order) statusChange(status string) {
// 	// function attached to struct

// 	o.status = status // struct derefrence automatically
// }

// // for modification using pointer is required but getting will work fine for both

// func (o order) getAmount() float32 {
// 	return o.amount
// }

// func main() {

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
// myOrder := order{
// id: "12",
// status: "delivered",
// amount: 100,
// createAt: time.Now(),
// }

// fmt.Println(myOrder)

// }

// creating constructor --> conventional name ( new )
// func New
// func newOrder(id string, amount float32, status string) *order { // work of creation abstract
// 	//  initial setup goes here
// 	myOrder := order{ // no need of dereferencing here because the struct is already doing it automatically
// 		id: id,
// 		status: status,
// 		amount: amount,
// 	}

// 	return &myOrder
// }

// func main() {

// 	// myOrder := newOrder("1", 50.50, "received")
// 	// fmt.Println(myOrder, myOrder.amount)

// }

// func main() {
// 	// object literal like js
// 	language := struct {
// 		name string
// 		isGood bool
// 	} {"goland", true}

// 	fmt.Println(language)
// }

/// STRUCT EMBEDDING
// struct inside struct

type customer struct {
	name string 
	phone string 
}

type order struct {
	id string 
	amount float32
	status string 
	createdAt time.Time 
	customer // other struct
}

func main() {
	// newOrder := order {
	// 	id: "1",
	// 	amount: 70,
	// 	status:  "received",
	// }

	// fmt.Println(newOrder)
	// fmt.Println(newOrder.customer)

	// newCustomer := customer {
	// 	name: "jay",
	// 	phone: "1234567890",
	// }
	// newOrder := order {
	// 	id: "1",
	// 	amount: 70,
	// 	status:  "received",
	// 	customer: newCustomer,
	// }


	newOrder := order {
		id: "1",
		amount: 70,
		status:  "received",
		customer: customer{
			name: "jay",
			phone: "1234567890",
		},
	}

	// fmt.Println(newOrder)
	// fmt.Println(newOrder.customer)

	newOrder.customer.name = "brave"
	fmt.Println(newOrder)
}

// through embeddings we can do --> composition, inheritance etc 