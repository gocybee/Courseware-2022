package main

import "main/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
	boot.ServerSetup()
}
