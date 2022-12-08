package main

import (
	"fmt"
	"net/http"
)

func testHelloWorld() {
	fmt.Println("hello world!")
}

func testHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	panic(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func main() {
	testHttp()
}
