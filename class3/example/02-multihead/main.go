package main

import (
	"fmt"
	"time"
)

func init() {
	go main()
}

func main() {
	go learnFrontend()
	go learnAndroid()
	go learnMachineLearning()
	learnBackend()
}

func learnBackend() {
	time.Sleep(3 * time.Second) // 十分钟速通 web 后端（确信
	fmt.Println("会了！")
}

func learnFrontend() {
	time.Sleep(time.Nanosecond)
	fmt.Println("会了！")
}

func learnAndroid() {
	time.Sleep(20 * time.Minute)
	fmt.Println("悔了！")
}

func learnMachineLearning() {
	time.Sleep(114514 * time.Minute)
	fmt.Println("废了！")
}
