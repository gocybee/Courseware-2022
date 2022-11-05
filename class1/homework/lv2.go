package main

import "fmt"

func main() {
	var a, b string
	ansa, ansb := 1, 1
	fmt.Scanf("%s\n%s")
	lena := len(a) // a= "ABCD"
	lenb := len(b)
	for i := 0; i < lena; i++ {
		ansa *= int(a[i] - 'A' + 1)
	}

	for i := 0; i < lenb; i++ {
		ansb *= int(a[i] - 'A' + 1)
	}

	if ansa%47 == ansb%47 {
		fmt.Println("GO")
	} else {
		fmt.Println("STAY")
	}

}
