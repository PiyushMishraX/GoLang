package main

import (
	"fmt"
)

// interfaces

// type payment struct {}

// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw := razorpay{}
// 	// // razorpayPaymentGw.pay(amount)
// 	// stripePaymentGw := stripe{}
// 	// stripePaymentGw.pay(amount) // we are modifying our code // solid principle
// 	// Open close principle violated
// 	// the classes should be open for extension but close for modification
// }

// type payment struct {
// 	// gateway stripe
// 	gateway razorpay

// }

// func (p payment) makePayment(amount float32) {
// 	p.gateway.pay(amount)
// }

// type razorpay struct {}

// func (r razorpay) pay(amount float32){
// 	// logic to make payment
// 	fmt.Println("Make payment using rozorpay", amount)
// }

// // if we have to use stripe for payment

// type stripe struct {}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

// type fakepayment struct { // for testing we create fake payment gateway not use the real one
// }
// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing purposes")
// }

// func main() {
// newPayment := payment{} // the struct values are not compulsory , so that might cause some type of mismatch use constructor function to make values compulsory to be defined
// newPayment.makePayment(100)

// stripePaymentGw := stripe{}
// newPayment := payment{
// 	gateway: stripePaymentGw,
// }

// razorpayPaymentGw := razorpay{}
// newPayment := payment{
// 	gateway: razorpayPaymentGw,
// }

// fakeGw := fakepayment{}
// newPayment := payment{
// 	gateway: fakeGw, // can not use without changing in payment struct // the solution for it is interfaces
// }

// 	newPayment.makePayment(100)
// }

// logger, storer
type paymenter interface {
	// any struct implementing it should have the method like below
	pay(amount float32) // if returing write "type" 
	// the function implementing it should have smae named and parameters as this else the error will show  ( same method signature required)
	// in go the implement keyword not reuqired it emplicitily find the interface itself

	refund(amount float32, account string) /// should have both method too in the struct
}	


type payment struct {
	// gateway razorpay // no concreate implementation name
	gateway paymenter //gateway should be of paymenter type

}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}

// func (r razorpay) pay2(amount float32){
func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Make payment using rozorpay", amount)
}

// type stripe struct {}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

// type fakepayment struct { }
// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing purposes")
// }

type paypal struct{}

func(p paypal) pay(amount float32){
	fmt.Println("Made payment using paypal of amount- ", amount)
}
func(p paypal) refund(amount float32, account string){
	fmt.Println("Made payment using paypal of amount- ", amount)
}

// open close principal fixed only etension no modification
func main() {

	// fakeGw := fakepayment{}
	// razorpayPaymentGw:= razorpay{}
	paypalGw := paypal{}

	newPayment := payment {
		// gateway: fakeGw,
		// gateway: razorpayPaymentGw,
		gateway: paypalGw,
	}
	newPayment.makePayment(100)

}