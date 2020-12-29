package models

import (
	"fmt"
	"go-gin/global"
)

type User struct {
	Id       int    `gorm:"Column:id" json:"id"`
	Username string `gorm:"Column:username" json:"username"`
	Password string `gorm:"Column:password" json:"password"`
}

func (*User) TableName() string {
	return "gin_users"
}

func (this *User) Create(user *User) (*User, error) {
	fmt.Println(user)
	fmt.Println(global.DB.Statement)
	if err := global.DB.Save(user).Error; err != nil {
		fmt.Println(err)
		return user, err
	}
	//if err := global.DB.Select("username", "password").Create(&user).Error; err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	return user, nil
}
