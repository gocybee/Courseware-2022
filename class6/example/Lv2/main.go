package main

type User struct {
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
	Question string `gorm:"question"`
	Answer   string `gorm:"answer"`
}
