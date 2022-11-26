package model

import "time"

type UserSubject struct {
	Id         int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Username   string    `gorm:"column:username" json:"username" form:"username" db:"username"`
	Password   string    `gorm:"column:password" json:"password" form:"password" db:"password"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
