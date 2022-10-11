package main

import (
	_ "Lv2/dao"
	"Lv2/global"
	"fmt"
)

func main() {
	var user = global.User{
		Name:     "TR",
		Account:  "123456",
		Password: "654321",
		Question: "Who are you?",
		Answer:   "啊吧啊吧啊吧",
	}

	//增--创建信息
	if err := global.DB.Model(global.User{}).Create(&user).Error; err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
	}
	MsgPrint()

	//改--更改信息
	if err := global.DB.Model(global.User{}).Update("password", "10086").Error; err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
	}
	MsgPrint()

	//查--检索信息
	var t global.User
	if err := global.DB.Model(global.User{}).Where("password=?", "10086").Find(&t).Error; err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
	}
	MsgPrint()

	//删-删除信息
	if err := global.DB.Model(global.User{}).Delete(&global.User{}).Error; err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
	}
	MsgPrint()
}

func MsgPrint() {
	var t global.User
	if err := global.DB.Model(global.User{}).First(&t).Error; err != nil {
		fmt.Println("[WARNING] an error occurred:", err)
		return
	}
	fmt.Println(t)
}
