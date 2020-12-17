package models

type User struct {
	UserId   int    `gorm:"Column:user_id" json:"user_id"`
	Username string `gorm:"Column:username" json:"username"`
	Password string `gorm:"Column:password" json:"password"`
}

func (*User) TableName() string {
	return "gin_users"
}
