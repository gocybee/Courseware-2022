package main

import "fmt"

func main() {
	fmt.Println(0)
	fmt.Println(10)
	for i := 1; i <= 5; i++ {
		fmt.Println(i*10 + (10 - i))
	}

	for i := 3; i <= 3; i++ { // 打出的数的长度
		for j := 1; j <= 9; j++ { // 打出的非0数
			for k := 1; k <= i; k++ { // 打出的第一个非0数的位置
				for l := 1; l <= i; l++ { // 打出的第二个非0数的位置
					if k != 1 && l != 1 {
						continue
					}
					for p := 1; p <= i; p++ { // 当前打印的数的位置
						if k == p {
							// 打印第一个非0数
							fmt.Print(j)
						} else if l == p {
							// 打印第二个非0数
							fmt.Print(10 - j)
						} else {
							fmt.Print(0)
						}
					}
					fmt.Println()
					for p := 1; p <= i; p++ {
						if k == p {
							// 打印第一个非0数
							fmt.Print(10 - j)
						} else if l == p {
							// 打印第二个非0数
							fmt.Print(j)
						} else {
							fmt.Print(0)
						}
					}
					fmt.Println()
				}
			}
			for k := 1; k <= i; k++ {
				if k == 1 {
					fmt.Print(j)
				} else {
					fmt.Print(0)
				}
			}
			fmt.Println()
		}
	}
}
