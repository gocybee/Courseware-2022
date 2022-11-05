package global

import "github.com/jinzhu/gorm"

// User 描述用户的基本信息结构体
type User struct {
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
	Question string `gorm:"question"`
	Answer   string `gorm:"answer"`
}

var DB *gorm.DB
