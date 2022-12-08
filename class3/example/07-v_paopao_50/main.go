package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	Wallet = 0 // 一贫如洗的泡泡
	lock   = &sync.Mutex{}
)

func main() {
	for i := 0; i < 10_000; i++ { // 泡泡成功骗取到了 1w 个人的同情
		go vPaopao50()
	}
	time.Sleep(2 * time.Second) // 可恶的泡泡竟然睡起了大觉
	fmt.Println("泡泡现在有", Wallet, "元", "Num of goroutine", runtime.NumGoroutine())
	// 睡醒的泡泡真的可以获得他乞讨到的 1w * 50 = 50w 元么？
}

func vPaopao50() {
	lock.Lock() // 阻塞
	defer lock.Unlock()
	Wallet += 50
}
