package dao

import (
	"Lv2/global"
	"fmt"
	"github.com/jinzhu/gorm"
)

// init 链接数据库并进行相关的初始化
func init() {
	// 链接数据库
	dsn := "user:@tcp(127.0.0.1:3306)/qff_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
	}

	// 判断表格是否存在，否 则创建一个表格
	if !db.HasTable(&global.User{}) {
		err = db.CreateTable(&global.User{}).Error
		if err != nil {
			fmt.Println("[WARNING] an error occurred:", err)
		}
	}

	//将数据复制给全局变量
	global.DB = db
}
