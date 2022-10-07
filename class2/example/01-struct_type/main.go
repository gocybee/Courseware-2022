package main

import (
	"fmt"
	"reflect"
)

type wxjj int

type wxgg = int

func main() {
	var a wxjj

	var b wxgg

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
	//rfTypeOf(a)
	//rfTypeOf(b)

	//TypeOf(a)
	//TypeOf(b)
}

func rfTypeOf(data interface{}) {
	of := reflect.TypeOf(data)
	fmt.Println(of)
}

func TypeOf(data interface{}) {
	switch data.(type) {
	case wxgg:
		fmt.Println("Type is int")
	case wxjj:
		fmt.Println("Type is wxjj")
	default:
		fmt.Println("Type Not Found")
	}
}
