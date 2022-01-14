package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`    // 设置用户名(name)不为空
	Phone    string `gorm:"varchar(110);not null;unique"` // 设置电话号码(phone)唯一并且不为空
	Password string `gorm:"size:255;not null"`            // 设置字段大小为255
}
