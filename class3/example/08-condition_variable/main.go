package main

import (
	"fmt"
	"sync"
)

var n, count = 5, 0

func main() {
	cv := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 8; i++ {
		go Producer(cv)
		go Consumer(cv)
	}

	<-make(chan struct{})
}

func Producer(cond *sync.Cond) {
	for {
		cond.L.Lock()
		for !(count != n) { // to assume count != n
			fmt.Println("=== Produce start waiting...")
			cond.Wait() // if count == n, continue waiting.
		}
		count++
		fmt.Println("produce a num, now:", count)
		cond.L.Unlock()
		cond.Broadcast() // to avoid waking a producer up and cause dead lock
	}

}

func Consumer(cond *sync.Cond) {
	for {
		cond.L.Lock()
		for !(count != 0) { // to assume count != 0
			fmt.Println("=== Consumer start waiting...")
			cond.Wait() // if count == 0, continue waiting.
		}
		count--
		fmt.Println("consume a num, now:", count)
		cond.L.Unlock()
		cond.Broadcast() // to avoid waking a consumer up and cause dead lock
	}
}
