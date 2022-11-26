package service

import "main/app/internal/service/user"

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
