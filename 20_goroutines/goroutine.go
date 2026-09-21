// goroutines

package main

import (
	"fmt"
	"time"
)


func task(id int) {
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




	for i := 0; i<=10; i++ {
		// go task(i)

		// inline fn
		// go func () {
		// 	fmt.Println(i) // i closure
		// } ()

		// better practice , receive the i even if closure
		go func (i int) {
			fmt.Println(i)
		} (i)
	}


	// go routine --> we can write background workers ,do cpu intensive work concurrently

	time.Sleep(time.Second * 2) // 2 second main holded 

}