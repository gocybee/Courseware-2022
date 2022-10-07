package main

import "fmt"

type WxGG struct {
	Name  string
	Age   int
	Books []Book
}

type Book struct {
	Name string
}

func (w WxGG) PrintName() {
	fmt.Println(w.Name)
}
func (w WxGG) PrintAge() {
	fmt.Println(w.Age)
}

func (w WxGG) PrintBook() {
	fmt.Println(w.Books)
}

func (b Book) PrintBookName() {
	fmt.Println(b.Name)
}

func main() {
	a := WxGG{
		Name: "wxGG",
		Age:  18,
		Books: []Book{
			{
				"Go圣经",
			},
			{
				"大话数据结构",
			},
		},
	}
	a.PrintName()
	a.PrintAge()
	a.PrintBook()
	for i := 0; i < len(a.Books); i++ {
		a.Books[i].PrintBookName()
	}
}
