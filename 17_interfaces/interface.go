package main

import "fmt"

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

type payment struct {
	// gateway stripe
	gateway razorpay

}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Make payment using rozorpay", amount)
}

// if we have to use stripe for payment

type stripe struct {}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

func main() {
	// newPayment := payment{} // the struct values are not compulsory , so that might cause some type of mismatch use constructor function to make values compulsory to be defined
	// newPayment.makePayment(100)


	// stripePaymentGw := stripe{}
	// newPayment := payment{
	// 	gateway: stripePaymentGw,
	// }
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: stripe(razorpayPaymentGw),
	}
	newPayment.makePayment(100)
}