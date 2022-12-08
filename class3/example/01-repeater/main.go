package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = &sync.Mutex{}
	lock2 = &sync.Mutex{}
)

func main() {
	go repeater1()
	go repeater2()
	go repeater2()
	go repeater2()
	go repeater1()
	go repeater1()
	go repeater1()
	go repeater2()
	<-make(chan struct{})
	// go repeater()
}

func repeater1() {
	for {
		time.Sleep(time.Second)
		lock1.Lock()
		lock2.Lock()
		fmt.Println("over.") // 这个是原子性？？
		// ....xx
		lock1.Unlock()
		lock2.Unlock()
	}
}

func repeater2() {
	for {
		time.Sleep(time.Second)
		lock2.Lock()
		lock1.Lock()
		fmt.Println("over.") // 这个是原子性？？
		// xx..
		lock2.Unlock()
		lock1.Unlock()
	}
}
