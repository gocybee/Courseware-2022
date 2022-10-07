package main

import "fmt"

func main() {
	var a *string
	*a = "无香的一刀"
	fmt.Println(*a)

	var b map[string]string
	b["袁神"] = "YYDS"
	fmt.Println(b)
}
