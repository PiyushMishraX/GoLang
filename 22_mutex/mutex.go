package main

import (
	"fmt"
	"sync"
)

// type post struct {
// 	views int
// }

// func (p *post) inc() {
// 	p.views += 1
// }

// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer wg.Done() // defer run at last
// 	p.views += 1
// }


// mutex
type post struct {
	views int

	// mu name convension
	mu sync.Mutex 
}

func (p *post) inc(wg *sync.WaitGroup) {
	// defer wg.Done() 
	defer func ()  {
		p.mu.Unlock()
		wg.Done()
	} ()
	
	p.mu.Lock()  // wait till the first goroutine modifies then run
	// downside of mutex is above , becuase we wan't to run goroutines concurrently but here is have to serial wise , which puts away concurrency
	// so better thing is put put lock only on the line needed not the other code above or below  like databases check etc in lines above // do not lock whle fn  // only in modification lines to reduce bottleneck

	p.views += 1

	// p.mu.Unlock() // better to put in defer // because real code might throw error so the ulock can not be reached so better to put in defer

}

func main() {
	// race condition
	// multi resources modifies same resource
	// than modification not atomic

	// myPost := post{views: 0}


	// myPost.inc()
	// myPost.inc()
	// many times in real operations , these operations are concurrent( both at same )

	// for i := 0; i < 100; i++ {
	// 	// myPost.inc() // here no problem		
	// 	go myPost.inc()  // concurrent
	// }

	// var wg sync.WaitGroup

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)	

	// 	go myPost.inc(&wg) // some processes run concurrently so for 2 processes one increment occurs
	// }

	// wg.Wait() // wait till all wait groups end // wg again = 0

	// fmt.Println(myPost.views)




	myPost := post{views: 0}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)	

		go myPost.inc(&wg) // some processes run concurrently so for 2 processes one increment occurs
		// fix mutex ,, can make globally like wg but better convension is one per struct
		// mutex for holding resource // locking resources
	}

	wg.Wait() 

	fmt.Println(myPost.views)

}