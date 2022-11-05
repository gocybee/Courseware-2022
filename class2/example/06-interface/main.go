package main

import (
	"fmt"
)

type Sayer interface {
	Say()
}

type dog struct {
}

type cat struct {
}

func (d dog) Say() {
	fmt.Println("汪汪汪")
}

func (c cat) Say() {
	fmt.Println("喵喵喵")
}
func main() {
	//var x Sayer
	//a := cat{}
	//b := dog{}
	//x = a
	//x.Say()
	//x = b
	//x.Say()

	var x interface{}
	s := "YuanShen"
	x = s
	v, ok := x.(string)
	if ok {
		fmt.Println("类型断言成功")
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
