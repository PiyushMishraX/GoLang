// goroutines

package main

import (
	"fmt"
	"sync"
	// "time"
)


// func task(id int) {
// 	fmt.Println("Doing task", id)
// }

func task(id int, w *sync.WaitGroup) { // pass pointer

	defer w.Done()
	// subtract 1 // react use effect cleaning  // this runs after function completes 
	// it runs after fn/goroutine / lightweight thread ends 


	fmt.Println("Doing task", id)
}

func main() {

	// for i := 0; i<=10; i++ {
	// 	task(i)
	// } // have a order 0 to 10

	// but we want to run them parallely
	// for i := 0; i<=10; i++ {
	// 	go task(i) // runs in lightweight thread(goroutine) concorrently
	// 	// schedular scheules them seperately without blocking while the main functions ends seperately( exist)
	// }  // very fast compared

	// go fn()




	// for i := 0; i<=10; i++ {
	// 	// go task(i)

	// 	// inline fn
	// 	// go func () {
	// 	// 	fmt.Println(i) // i closure
	// 	// } ()

	// 	// better practice , receive the i even if closure
	// 	go func (i int) {
	// 		fmt.Println(i)
	// 	} (i)
	// }


	// go routine --> we can write background workers ,do cpu intensive work concurrently


	// time.Sleep(time.Second * 2) // 2 second main holded 


	// ------------------	WAIT GROUP -------

	// go routine have problem/s --> wait groups
   // we don't know how much time will the goroutine execcution take so we can't just use 2 second hold 
   // so we use wait group

   var wg sync.WaitGroup

   for i := 0; i<=10; i++ {
		wg.Add(1) // add 1 // in each goroutine // light weight thread
		go task(i, &wg)  // send add of wg for proper real address passing

		// for inline fn we do not need to pass 
		// we can directly add and done in it
	}

	// wg --> add , done , wait

	wg.Wait() // so the code runs till wg again becomes 0 // by deffault 0 too // second 0 means everything done



}