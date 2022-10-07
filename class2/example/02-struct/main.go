package main

import (
	"fmt"
)

type WxGG struct {
	Name string
	Age  int
}

func main() {

	//最常见的方式
	a := WxGG{
		Name: "wxgg1",
		Age:  18,
	}

	var b WxGG
	b.Name = "wxgg2"
	b.Age = 18

	//var WxJJ struct {
	//	Name string
	//	Age  int
	//}
	//
	//WxJJ.Name = "wxjj1"
	//WxJJ.Age = 18

	WxJJ := struct {
		Name string
		Age  int
	}{
		"wxjj2",
		18,
	}

	//  类比
	//	type  yxh int
	//	god :=yxh(55)

	c := NewWxGG("wxgg tql", 18)

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", WxJJ)
	fmt.Printf("%#v\n", c)

}

func NewWxGG(name string, age int) *WxGG {
	return &WxGG{
		Name: name,
		Age:  age,
	}
}
