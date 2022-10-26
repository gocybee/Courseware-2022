package main

import (
	"fmt"
	"time"
)

func main() {
	repeater()
}

func repeater() {
	for {
		time.Sleep(time.Second)
		fmt.Println("over.")
	}
}
