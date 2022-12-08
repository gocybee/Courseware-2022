package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018

	//var p *string
	//fmt.Println(p)
	//fmt.Printf("p的值是%v\n", p)
	//if p != nil {
	//	fmt.Println("非空")
	//} else {
	//	fmt.Println("空值")
	//}

}
